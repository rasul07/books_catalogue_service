package postgres

import (
	"book_catalog_service/models"
	"book_catalog_service/storage/repo"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type bookService struct {
	db *sqlx.DB
}

func NewBookRepo(db *sqlx.DB) repo.BooksRepoI {
	return bookService{db: db}
}

func (c bookService) Create(req models.BookCreate) (err error) {
	query := `
	INSERT INTO book(
		guid,
		name,
		author,
		category_id,
		description,
		pages,
		year
	)
	VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7
	)`

	bookID := uuid.New()
	category, err := uuid.Parse(req.Category)
	if err != nil {
		return
	}
	year, err := time.Parse("2006-01-02", req.Year)
	if err != nil {
		return
	}
	_, err = c.db.Exec(
		query,
		bookID,
		req.Name,
		req.Author,
		category,
		req.Description,
		req.Pages,
		year,
	)

	if err != nil {
		return
	}

	return
}

func (c bookService) GetBookList(limit, page int32) (resp []models.Book) {
	var (
		offset = (page - 1) * limit
	)

	query := `
		SELECT
			count(1) over(),
			guid,
			name,
			author,
			category_id,
			description,
			pages,
			year
		FROM book
		limit $2 offset $3
	`

	rows, err := c.db.Query(query, limit, offset)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var book models.Book
		err = rows.Scan(
			&book.Count,
			&book.ID,
			&book.Name,
			&book.Author,
			&book.Category,
			&book.Description,
			&book.Pages,
			&book.Year,
		)

		if err != nil {
			return
		}

		resp = append(resp, book)
	}

	return
}

func (c bookService) GetBookById(BookID string) (resp models.Book, err error) {
	query := `
		SELECT
			guid,
			name,
			author,
			category_id,
			description,
			pages,
			year
		FROM book
		WHERE guid=$1
	`

	err = c.db.QueryRow(query, BookID).Scan(
		&resp.ID,
		&resp.Name,
		&resp.Author,
		&resp.Category,
		&resp.Description,
		&resp.Pages,
		&resp.Year,
	)

	if err != nil {
		return models.Book{}, err
	}

	return
}

func (c bookService) Update(req models.BookCreate) (err error) {
	query := `
	UPDATE book SET name = $1, author = $2, category = $3, description = $4, pages = $5, year = $6`

	category, err := uuid.Parse(req.Category)
	if err != nil {
		return
	}
	year, err := time.Parse("2006-01-02", req.Year)
	if err != nil {
		return
	}
	_, err = c.db.Exec(
		query,
		req.Name,
		req.Author,
		category,
		req.Description,
		req.Pages,
		year,
	)

	if err != nil {
		return
	}

	return
}