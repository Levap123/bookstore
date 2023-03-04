package book

import (
	"time"

	"github.com/Levap123/book_service/proto"
)

type Book struct {
	ID          string    `bson:"_id,omitempty"`
	Title       string    `bson:"title,omitempty"`
	Description string    `bson:"description,omitempty"`
	Image       string    `bson:"image,omitempty"`
	Pages       uint64    `bson:"pages,omitempty"`
	Author      string    `bson:"author,omitempty"`
	Genre       string    `bson:"genre,omitempty"`
	Publisher   string    `bson:"publisher,omitempty"`
	Binding     bool      `bson:"binding,omitempty"`
	Series      string    `bson:"series,omitempty"`
	Language    string    `bson:"language,omitempty"`
	CreatedAt   time.Time `bson:"created_at,omitempty"`
}

func NewBookFromCreateBookRequest(req *proto.CreateBookRequest) *Book {
	return &Book{
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		Pages:       req.Pages,
		Author:      req.Author,
		Genre:       req.Genre,
		Publisher:   req.Publisher,
		Binding:     req.Binding,
		Series:      req.Series,
		Language:    req.Language,
		CreatedAt:   time.Now(),
	}
}
