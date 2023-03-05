package dto

import "github.com/Levap123/api_gateway/proto"

type CreateBookDTO struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Image       string `json:"image,omitempty"`
	Pages       uint64 `json:"pages,omitempty"`
	Author      string `json:"author,omitempty"`
	Genre       string `json:"genre,omitempty"`
	Publisher   string `json:"publisher,omitempty"`
	Binding     bool   `json:"binding,omitempty"`
	Series      string `json:"series,omitempty"`
	Language    string `json:"language,omitempty"`
}

func FromDtoToRequest(dto *CreateBookDTO) *proto.BookInfo {
	return &proto.BookInfo{
		Title:       dto.Title,
		Description: dto.Description,
		Image:       dto.Image,
		Pages:       dto.Pages,
		Author:      dto.Author,
		Genre:       dto.Genre,
		Publisher:   dto.Publisher,
		Binding:     dto.Binding,
		Series:      dto.Series,
		Language:    dto.Language,
	}
}
