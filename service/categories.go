package service

import (
	"book_catalog_service/genproto/category"
	"book_catalog_service/pkg/helper"
	"book_catalog_service/pkg/logger"
	"book_catalog_service/storage"
	"context"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"
)

type categoryService struct {
	logger  logger.Logger
	storage storage.StorageI
	category.UnimplementedCategoryServiceServer
}

func NewCategoryService(log logger.Logger, db *sqlx.DB) *categoryService {
	return &categoryService{
		logger:  log,
		storage: storage.NewStoragePG(db),
	}
}

func (s *categoryService) Create(ctx context.Context, req *category.CategoryCreate) (resp *emptypb.Empty, err error) {
	resp = &emptypb.Empty{}

	err = s.storage.CategoriesRepo().Create(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while creating category ", req, codes.Internal)
	}

	return resp, nil
}

func (s *categoryService) GetCategoryById(ctx context.Context, req *category.GetCategoryByIdRequest) (resp *category.GetCategoryByIdResponse, err error) {
	s.logger.Info("Get category by id", logger.Any("req : ", req))
	category, err := s.storage.CategoriesRepo().GetCategoryById(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting category ", req, codes.Internal)
	}
	s.logger.Info("Get category by id : req", logger.Any("resp : ", resp))
	return category, nil
}

func (s *categoryService) GetCategories(ctx context.Context, req *category.GetAllCategoriesRequest) (resp *category.GetAllCategoriesResponse, err error) {
	categories, err := s.storage.CategoriesRepo().GetCategories(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while getting all categories ", req, codes.Internal)
	}

	return categories, nil
}

func (s *categoryService) Update(ctx context.Context, req *category.Category) (resp *emptypb.Empty, err error) {
	resp = &emptypb.Empty{}

	err = s.storage.CategoriesRepo().Update(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while updating category ", req, codes.Internal)
	}
	return resp, nil
}

func (s *categoryService) Delete(ctx context.Context, req *category.Category) (resp *emptypb.Empty, err error) {
	resp = &emptypb.Empty{}

	err = s.storage.CategoriesRepo().Delete(req)
	if err != nil {
		return nil, helper.HandleError(s.logger, err, "error while deleting category ", req, codes.Internal)
	}

	return resp, nil
}
