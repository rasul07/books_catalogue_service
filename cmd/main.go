package main

import (
	"fmt"

	"book_catalog_service/config"

	_ "github.com/jackc/pgx/v4/stdlib"

	// "book_catalog_service/pkg/logger"
	"book_catalog_service/storage"

	"github.com/jmoiron/sqlx"
)

func main() {

	cfg := config.Load()
	// logger := logger.New(cfg.LogLevel, "safia_catalogue_service")

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s", cfg.PostgresHost,
		cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase, "disable")
	db, err := sqlx.Connect("pgx", conStr)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(10)

	storageI := storage.NewStoragePG(db)

	// apiServer := api.New(&api.RouterOptions{
	// 	Cfg:     &cfg,
	// 	Storage: storageI,
	// })

	// err = apiServer.Run(cfg.HttpPort)
	// if err != nil {
	// 	panic(err)
	// }

}
