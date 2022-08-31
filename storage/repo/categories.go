package repo

import (
	"book_catalog_service/genproto/category"
)

type CategoriesRepoI interface {
	Create(req *category.CategoryCreate) (err error)
	Update(req *category.Category) (err error)
	GetCategories(*category.GetAllCategoriesRequest) (resp *category.GetAllCategoriesResponse, err error)
	GetCategoryById(*category.GetCategoryByIdRequest) (resp *category.GetCategoryByIdResponse, err error)
	Delete(*category.Category) (err error)

}
