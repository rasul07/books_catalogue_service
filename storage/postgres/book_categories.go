package postgres

import (
	"book_catalog_service/models"
	"book_catalog_service/storage/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type bookCategoryService struct {
	db *sqlx.DB
}

func NewBookCategoryRepo(db *sqlx.DB) repo.BookCategoriesRepoI {
	return bookCategoryService{db: db}
}

func (c bookCategoryService) Create(req models.CategoryCreate) (err error) {
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

func (c bookCategoryService) GetCategories(limit, page int32) (resp []models.Category, err error) {
	var (
		offset = (page - 1) * limit
	)

	query := `
		SELECT
			guid,
			name
		FROM book_category
		limit $1 offset $2;
	`
	rows, err := c.db.Query(query, limit, offset)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var category models.Category
		err = rows.Scan(
			&category.ID,
			&category.Name,
		)

		if err != nil {
			return
		}

		resp = append(resp, category)
	}

	return
}

func (c bookCategoryService) GetCategoryById(CategoryID string) (resp models.Category, err error) {
	query := `
		SELECT
			guid,
			name
		FROM book_category
		WHERE guid=$1
	`

	err = c.db.QueryRow(query, CategoryID).Scan(
		&resp.ID,
		&resp.Name,
	)

	if err != nil {
		return models.Category{}, err
	}

	return
}

func (c bookCategoryService) Update(req models.CategoryCreate) (err error) {
	query := `UPDATE book_category SET name = $1`

	_, err = c.db.Exec(
		query,
		req.Name,
	)

	if err != nil {
		return
	}

	return
}