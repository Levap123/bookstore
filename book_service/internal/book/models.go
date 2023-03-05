package book

import (
	"time"

	"github.com/Levap123/book_service/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Book struct {
	ID          string    `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string    `bson:"title,omitempty" json:"title,omitempty"`
	Description string    `bson:"description,omitempty" json:"description,omitempty"`
	Image       string    `bson:"image,omitempty" json:"image,omitempty"`
	Pages       uint64    `bson:"pages,omitempty" json:"pages,omitempty"`
	Author      string    `bson:"author,omitempty" json:"author,omitempty"`
	Genre       string    `bson:"genre,omitempty" json:"genre,omitempty"`
	Publisher   string    `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Binding     bool      `bson:"binding,omitempty" json:"binding,omitempty"`
	Series      string    `bson:"series,omitempty" json:"series,omitempty"`
	Language    string    `bson:"language,omitempty" json:"language,omitempty"`
	AddedAt     time.Time `bson:"created_at,omitempty" json:"added_at,omitempty"`
}

func NewBookFromCreateBookRequest(req *proto.BookInfo) *Book {
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
		AddedAt:     time.Now(),
	}
}

func NewBookResponseFromBook(book *Book) *proto.BookInfo {
	return &proto.BookInfo{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		Image:       book.Image,
		Pages:       book.Pages,
		Author:      book.Author,
		Genre:       book.Genre,
		Publisher:   book.Publisher,
		Binding:     book.Binding,
		Series:      book.Series,
		Language:    book.Language,
		AddedAt:     timestamppb.New(book.AddedAt),
	}
}
