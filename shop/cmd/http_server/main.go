package main

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"gitlab.com/tirava/shop/internal/config"
	http "gitlab.com/tirava/shop/internal/http_server"
	"gitlab.com/tirava/shop/internal/storages"
)

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger()

	cfg, err := config.NewDBConfig()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := storages.NewGormDB("postgres", dsn)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	defer db.Close()

	log.Info().Msgf("Connected to: %s", db.Dialect().GetName())

	srv := http.NewServer(":8000", db)
	if err := srv.StartServer(); err != nil {
		log.Fatal().Msg(err.Error())
	}
}
