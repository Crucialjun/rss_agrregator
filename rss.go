package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RssFeed struct {

	Channel struct {
		Title       string `xml:"title"`
		Description string `xml:"description"`
		Link        string `xml:"link"`
		Language    string `xml:"language"`
		Items       []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
}

func RssUrlToFeed(url string) (RssFeed, error) {
	httpClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return RssFeed{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RssFeed{}, err
	}

	rssFeed := RssFeed{}

	if err := xml.Unmarshal(data, &rssFeed); err != nil {
		return RssFeed{}, err
	}

	return rssFeed, nil
}
	

	

	
