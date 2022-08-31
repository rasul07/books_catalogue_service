package models

type BookCreate struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Pages       int32  `json:"pages"`
	Year        string `json:"year"`
}

type Book struct {
	ID          string `json:"guid"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Pages       int32  `json:"pages"`
	Year        string `json:"year"`
	Count       int32  `json:"count" swaggerignore:"true"`
}

type BookList struct {
	Books []Book `json:"books"`
	Count int32  `json:"count"`
	Limit int32  `json:"limit"`
	Page  int32  `json:"page"`
}
