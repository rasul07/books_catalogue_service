package storage

import (
	"book_catalog_service/storage/postgres"
	"book_catalog_service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	BookRepo() repo.BooksRepoI
	CategoriesRepo() repo.CategoriesRepoI
}

type storagePG struct {
	db       *sqlx.DB
	book     repo.BooksRepoI
	category repo.CategoriesRepoI
}

func NewStoragePG(db *sqlx.DB) StorageI {
	return storagePG{
		db:       db,
		book:     postgres.NewBookRepo(db),
		category: postgres.NewCategoryRepo(db),
	}
}

func (s storagePG) BookRepo() repo.BooksRepoI {
	return s.book
}

func (s storagePG) CategoriesRepo() repo.CategoriesRepoI {
	return s.category
}
