package main

import (
	"fmt"
	"net"

	"book_catalog_service/config"
	"book_catalog_service/genproto/book"
	"book_catalog_service/genproto/category"

	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"book_catalog_service/pkg/logger"
	"book_catalog_service/service"

	"github.com/jmoiron/sqlx"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.Environment, "catalogue_service")
	defer logger.Cleanup(log)

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	db, err := sqlx.Connect("pgx", conStr)
	if err != nil {
		log.Error("error while connecting database", logger.Error(err))
		return
	}

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Error("error while listening: %v", logger.Error(err))
		return
	}

	bookService := service.NewBookService(log, db)
	categoryService := service.NewCategoryService(log, db)

	s := grpc.NewServer()
	reflection.Register(s)

	book.RegisterBookServiceServer(s, bookService)
	category.RegisterCategoryServiceServer(s, categoryService)

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Error("error while listening: %v", logger.Error(err))
	}
}
