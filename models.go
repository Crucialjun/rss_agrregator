package main

import (
	"time"

	"github.com/crucialjun/rss_aggregator/internal/database"
	"github.com/google/uuid"
)




type User struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

}

func databaseUserToUser (dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		Email:     dbUser.Email,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}