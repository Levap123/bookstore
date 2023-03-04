package entity

type Book struct {
	ID          string `json:"id,omitempty"`
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
