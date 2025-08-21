package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/crucialjun/rss_aggregator/internal/database"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequests time.Duration,
) {
	log.Printf("Starting RSS feed scraping on %s with %d goroutines, waiting %v between requests", db, concurrency, timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	for ; ; <-ticker.C {

		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error fetching feeds: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)
			go func(f database.Feed) {
				defer wg.Done()
				log.Printf("Fetching feed: %s", f.Url)

				_, err := db.UpdateFeedLastFetched(context.Background(), f.ID)

				if err != nil {
					log.Printf("Error updating last fetched time for feed %d: %v", f.ID, err)
					return
				}

				rssFeed, err := RssUrlToFeed(f.Url)
				if err != nil {
					log.Printf("Error fetching RSS feed %s: %v", f.Url, err)
					return
				}

				for _, item := range rssFeed.Channel.Items {

					t, err := time.Parse(time.RFC1123Z, item.PubDate)
					if err != nil {
						log.Printf("Error parsing published date for item %s: %v", item.Title, err)
						continue
					}

					_, errr := db.CreatePost(context.Background(), database.CreatePostParams{
						FeedID:      f.ID,
						Title:       item.Title,
						Link:        item.Link,
						PublishedAt: t,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
						Url:         item.Link,
					})
					if errr != nil {
						log.Printf("Error creating post for feed %d: %v", f.ID, err)
						return
					}
				}

			}(feed)
		}
		wg.Wait()

		log.Printf("Completed fetching %d feeds, waiting for next cycle", len(feeds))

	}

}
