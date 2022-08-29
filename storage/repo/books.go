package repo

import (
	"book_catalog_service/models"
)

type BooksRepoI interface {
	Create(req models.BookCreate) (err error)
	Update(req models.BookCreate) (err error)
	GetBookList(limit, page int32) (resp []models.Book)
	GetBookById(BookID string) (resp models.Book, err error)
}
