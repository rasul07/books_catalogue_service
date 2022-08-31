package service

import (
	"book_catalog_service/genproto/book"
	"book_catalog_service/pkg/helper"
	"book_catalog_service/pkg/logger"
	"book_catalog_service/storage"
	"context"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type bookService struct {
	logger  logger.Logger
	storage storage.StorageI
	book.UnimplementedBookServiceServer
}

func NewBookService(log logger.Logger, db *sqlx.DB) *bookService {
	return &bookService{
		logger:  log,
		storage: storage.NewStoragePG(db),
	}
}

func (s *bookService) Create(ctx context.Context, req *book.BookCreate) (resp *emptypb.Empty, err error) {
	resp = &emptypb.Empty{}
	err = s.storage.BookRepo().Create(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating book ", req, codes.Internal)
	}

	return resp, nil
}

func (s *bookService) GetBookById(ctx context.Context, req *book.GetBookByIdRequest) (resp *book.GetBookByIdResponse, err error) {
	book, err := s.storage.BookRepo().GetBookById(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting book ", req, codes.Internal)
	}

	return book, nil
}

func (s *bookService) GetBooks(ctx context.Context, req *book.GetAllBooksRequest) (*book.GetAllBooksResponse, error) {
	books, err := s.storage.BookRepo().GetBookList(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all books ", req, codes.Internal)
	}

	return books, nil
}

func (s *bookService) Update(ctx context.Context, req *book.Book) (resp *emptypb.Empty, err error) {
	resp = &emptypb.Empty{}
	_, err = s.storage.BookRepo().Update(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating book ", req, codes.Internal)
	}
	return resp, err
}

func (s *bookService) Delete(ctx context.Context, req *book.Book) (resp *emptypb.Empty, err error) {
	resp = &emptypb.Empty{}
	err = s.storage.BookRepo().Delete(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while deleting book ", req, codes.Internal)
	}
	return resp, nil
}
