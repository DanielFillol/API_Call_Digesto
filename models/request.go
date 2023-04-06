package models

type BodyRequest struct {
	Document string     `json:"document,omitempty"`
	Pages    Pagination `json:"pagination,omitempty"`
}

type Pagination struct {
	Size   int    `json:"size,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}
