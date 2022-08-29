package repo

import (
	"book_catalog_service/models"
)

type BookCategoriesRepoI interface {
	Create(req models.CategoryCreate) (err error)
	Update(req models.CategoryCreate) (err error)
	GetCategories(limit, page int32) (resp []models.Category, err error)
	GetCategoryById(CategoryID string) (resp models.Category, err error)
}
