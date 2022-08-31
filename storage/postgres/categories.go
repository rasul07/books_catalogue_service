package postgres

import (
	"book_catalog_service/genproto/category"
	"book_catalog_service/storage/repo"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type bookCategoryService struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) repo.CategoriesRepoI {
	return bookCategoryService{db: db}
}

func (c bookCategoryService) Create(req *category.CategoryCreate) (err error) {
	query := `
	INSERT INTO book_category(
		guid,
		name
	)
	VALUES (
		$1,
		$2
	)`

	categoryID := uuid.New()

	_, err = c.db.Exec(
		query,
		categoryID,
		req.Name,
	)

	if err != nil {
		return
	}

	return
}

func (c bookCategoryService) GetCategories(req *category.GetAllCategoriesRequest) (resp *category.GetAllCategoriesResponse, err error) {
	var (
		categories  []*category.Category
		offset = (req.Page - 1) * req.Limit
	)

	query := `
		SELECT
			guid,
			name
		FROM book_category
		limit $1 offset $2;
	`
	rows, err := c.db.Query(query, req.Limit, offset)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var category category.Category
		err = rows.Scan(
			&category.ID,
			&category.Name,
		)

		if err != nil {
			return
		}

		categories = append(categories, &category)
	}

	return &category.GetAllCategoriesResponse{
		CategoryList: categories,
	}, nil
}

func (c bookCategoryService) GetCategoryById(req *category.GetCategoryByIdRequest) (resp *category.GetCategoryByIdResponse, err error) {
	var (
		categories category.Category
	)
	
	query := `
		SELECT
			guid,
			name
		FROM book_category
		WHERE guid=$1
	`
	fmt.Println("Guid =====>  ", req.BookId)
	err = c.db.QueryRow(query, req.BookId).Scan(
		&categories.ID,
		&categories.Name,
	)

	if err != nil {
		return nil, err
	}

	return &category.GetCategoryByIdResponse{
		Category: &categories,
	}, nil
}

func (c bookCategoryService) Update(req *category.Category) (err error) {
	query := `
	UPDATE 
		book_category 
	SET 
		name = $1
	WHERE 
		guid = $2`

	_, err = c.db.Exec(
		query,
		req.Name,
		req.ID,
	)

	if err != nil {
		return
	}

	return
}

func (c bookCategoryService) Delete(req *category.Category) (err error) {
	query := `
	DELETE FROM 
		book_category
	WHERE guid = $1`

	_, err = c.db.Exec(query, req.ID)
	if err != nil {
		return
	}

	return
}
