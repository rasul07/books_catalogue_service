package postgres

import (
	"book_catalog_service/genproto/book"
	"book_catalog_service/storage/repo"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"google.golang.org/protobuf/types/known/emptypb"
)

type bookService struct {
	db *sqlx.DB
}

func NewBookRepo(db *sqlx.DB) repo.BooksRepoI {
	return bookService{db: db}
}

func (c bookService) Create(req *book.BookCreate) (err error) {
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

func (c bookService) GetBookList(req *book.GetAllBooksRequest) (resp *book.GetAllBooksResponse, err error) {
	var (
		books    []*book.Book
		offset   = (req.Page - 1) * req.Limit
		category sql.NullString
	)

	query := `
		SELECT 
			guid,
			name,
			author,
			category_id,
			description,
			pages,
			to_char(year, 'DD.MM.YYYY')
		FROM book
		limit $1 offset $2
	`

	rows, err := c.db.Query(query, req.Limit, offset)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var book book.Book
		err = rows.Scan(
			&book.ID,
			&book.Name,
			&book.Author,
			&category,
			&book.Description,
			&book.Pages,
			&book.Year,
		)
		if err != nil {
			return
		}

		if category.Valid {
			book.Category = category.String
		}

		books = append(books, &book)
	}

	return &book.GetAllBooksResponse{
		BookList: books,
	}, nil
}

func (c bookService) GetBookById(req *book.GetBookByIdRequest) (resp *book.GetBookByIdResponse, err error) {
	var (
		books book.Book
	)

	query := `
		SELECT
			guid,
			name,
			author,
			category_id,
			description,
			pages,
			to_char(year, 'DD.MM.YYYY')
		FROM book
		WHERE guid=$1
	`

	err = c.db.QueryRow(query, req.BookId).Scan(
		&books.ID,
		&books.Name,
		&books.Author,
		&books.Category,
		&books.Description,
		&books.Pages,
		&books.Year,
	)
	if err != nil {
		return nil, err
	}

	return &book.GetBookByIdResponse{
		Book: &books,
	}, nil
}

func (c bookService) Update(req *book.Book) (resp *emptypb.Empty, err error) {
	tx, err := c.db.Begin()

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := `
	UPDATE 
		book 
	SET 
		name = $1, 
		author = $2, 
		category_id = $3, 
		description = $4,
		pages = $5,
		year = $6
	WHERE
		guid = $7`

	category, err := uuid.Parse(req.Category)
	if err != nil {
		return
	}
	year, err := time.Parse("2006-01-02", req.Year)
	if err != nil {
		return
	}
	_, err = tx.Exec(
		query,
		req.Name,
		req.Author,
		category,
		req.Description,
		req.Pages,
		year,
		req.ID,
	)
	if err != nil {
		return
	}

	return
}

func (c bookService) Delete(req *book.Book) (err error) {
	query := `
	DELETE FROM 
		book
	WHERE guid = $1`

	_, err = c.db.Exec(query, req.ID)
	if err != nil {
		return
	}

	return
}
