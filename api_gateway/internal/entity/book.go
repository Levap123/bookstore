package entity

import (
	"time"

	"github.com/Levap123/api_gateway/proto"
)

type Book struct {
	ID          string    `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Image       string    `json:"image,omitempty"`
	Pages       uint64    `json:"pages,omitempty"`
	Author      string    `json:"author,omitempty"`
	Genre       string    `json:"genre,omitempty"`
	Publisher   string    `json:"publisher,omitempty"`
	Binding     bool      `json:"binding,omitempty"`
	Series      string    `json:"series,omitempty"`
	Language    string    `json:"language,omitempty"`
	AddedAt     time.Time `json:"added_at,omitempty"`
}

func FromBookRequestToBook(req *proto.BookInfo) Book {
	return Book{
		ID:          req.ID,
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
		AddedAt:     req.AddedAt.AsTime(),
	}
}
