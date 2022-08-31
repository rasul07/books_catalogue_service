package models

type CategoryCreate struct {
	Name string `json:"category_name"`
}

type Category struct {
	ID   string `json:"guid"`
	Name string `json:"category_name"`
}

type BooksCategories struct {
	Categories []Category `json:"categories"`
	Limit      int32      `json:"limit"`
	Page       int32      `json:"page"`
}
