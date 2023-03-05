package entity

import (
	"time"

	"github.com/Levap123/api_gateway/proto"
)

type Book struct {
	ID          string
	Title       string
	Description string
	Image       string
	Pages       uint64
	Author      string
	Genre       string
	Publisher   string
	Binding     bool
	Series      string
	Language    string
	AddedAt     time.Time
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
