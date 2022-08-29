package storage

import (
	"book_catalog_service/storage/postgres"
	"book_catalog_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	BookRepo() repo.BooksRepoI
	BookCategoriesRepo() repo.BookCategoriesRepoI
}

type storagePG struct {
	db           *sqlx.DB
	book         repo.BooksRepoI
	bookCategory repo.BookCategoriesRepoI
}

func NewStoragePG(db *sqlx.DB) StorageI {
	return storagePG{
		db:           db,
		book:         postgres.NewBookRepo(db),
		bookCategory: postgres.NewBookCategoryRepo(db),
	}
}

func (s storagePG) BookRepo() repo.BooksRepoI {
	return s.book
}

func (s storagePG) BookCategoriesRepo() repo.BookCategoriesRepoI {
	return s.bookCategory
}
