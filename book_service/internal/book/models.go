package book

import "time"

type Book struct {
	ID          string    `bson:"id,omitempty"`
	Title       string    `bson:"title,omitempty"`
	Description string    `bson:"description,omitempty"`
	Image       string    `bson:"image,omitempty"`
	Pages       int       `bson:"pages,omitempty"`
	Author      string    `bson:"author,omitempty"`
	Genre       string    `bson:"genre,omitempty"`
	Publisher   string    `bson:"publisher,omitempty"`
	Binding     bool      `bson:"binding,omitempty"`
	Series      string    `bson:"series,omitempty"`
	Language    string    `bson:"language,omitempty"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
}
