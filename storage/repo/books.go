package repo

import (
	"book_catalog_service/genproto/book"

	"google.golang.org/protobuf/types/known/emptypb"
)

type BooksRepoI interface {
	Create(req *book.BookCreate) (err error)
	Update(req *book.Book) (resp *emptypb.Empty, err error)
	GetBookList(*book.GetAllBooksRequest) (resp *book.GetAllBooksResponse, err error)
	GetBookById(*book.GetBookByIdRequest) (resp *book.GetBookByIdResponse, err error)
	Delete(*book.Book) (err error)
}
