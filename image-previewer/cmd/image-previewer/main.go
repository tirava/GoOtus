/*
 * Project: Image Previewer
 * Created on 10.01.2020 13:20
 * Copyright (c) 2020 - Eugene Klimov
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gitlab.com/tirava/image-previewer/internal/helpers"

	"gitlab.com/tirava/image-previewer/internal/models"

	"gitlab.com/tirava/image-previewer/internal/http"
	"gitlab.com/tirava/image-previewer/internal/loggers"

	"gitlab.com/tirava/image-previewer/internal/configs"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
)

const previewConfigPath = "PREVIEWER_CONFIG_PATH"

func main() {
	fileName := filepath.Base(os.Args[0])
	flag.Usage = func() {
		fmt.Printf("Start http server:\n"+
			"%s [-config=configFile|inmemory]\n"+
			"[PREVIEWER_CONFIG_PATH=configFile|inmemory] %s\n",
			fileName, fileName)
		flag.PrintDefaults()
	}

	config := flag.String("config", "config.yml", "path to yaml config file or 'inmemory'")
	flag.Parse()

	if os.Getenv(previewConfigPath) != "" {
		*config = os.Getenv(previewConfigPath)
	}

	cfg, err := configs.NewConfig(*config)
	if err != nil {
		log.Fatal(err)
	}

	conf := cfg.GetConfig()

	logFile, err := os.OpenFile(conf.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error open log file '%s', error: %s", conf.LogFile, err)
	}
	defer logFile.Close()

	lg, err := loggers.NewLogger(conf.Logger, conf.LogLevel, logFile)
	if err != nil {
		log.Fatal(err)
	}

	prev, err := helpers.InitPreview(conf)
	if err != nil {
		log.Fatal(err)
	}

	opts := entities.ResizeOptions{
		Interpolation: conf.Interpolation,
	}

	if *config == "inmemory" {
		printConfig(conf)
	}

	log.Println("Logger started at mode:", conf.LogLevel)
	http.StartHTTPServer(lg, conf, *prev, opts)

	if err := prev.Close(); err != nil {
		log.Fatal("error close previewer:", err)
	}

	os.Exit(0)
}

func printConfig(conf models.Config) {
	log.Println("InMemory config:")
	log.Println("Logger:", conf.Logger)
	log.Println("LogFile:", conf.LogFile)
	log.Println("LogLevel:", conf.LogLevel)
	log.Println("ListenHTTP:", conf.ListenHTTP)
	log.Println("ListenPrometheus:", conf.ListenPrometheus)
	log.Println("Previewer:", conf.Previewer)
	log.Println("Interpolation:", "NearestNeighbor")
	log.Println("NoProxyHeaders:", conf.NoProxyHeaders)
	log.Println("ImageURLEncoder:", conf.ImageURLEncoder)
}
