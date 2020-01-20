/*
 * Project: Image Previewer
 * Created on 10.01.2020 13:20
 * Copyright (c) 2020 - Eugene Klimov
 */

package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"time"

	"gitlab.com/tirava/image-previewer/internal/models"

	"gitlab.com/tirava/image-previewer/internal/configs"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"

	"gitlab.com/tirava/image-previewer/internal/domain/preview"

	"github.com/google/uuid"
)

const (
	to   = 50
	fail = 1
)

func main() {
	fileName := filepath.Base(os.Args[0])
	flag.Usage = func() {
		fmt.Printf("Use config file: %s -config=configFile|inmemory\n"+
			"or instead you may use env variable PREVIEWER_CONFIG_PATH\n\n", fileName)
		fmt.Printf("Preview image tool w/o starting http server:\n"+
			"%s [-config=configFile] -preview -width=x -height=y -file=imageFile\n", fileName)
		flag.PrintDefaults()
	}

	config := flag.String("config", "config.yml", "path to yaml config file or 'inmemory'")
	previewMode := flag.Bool("preview", false, "use preview tool w/o starting http server")
	imageFile := flag.String("file", "", "path to image file")
	width := flag.Int("width", 0, "preview width")
	height := flag.Int("height", 0, "preview height")
	flag.Parse()

	cfg, err := configs.NewConfig(*config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.GetConfig())

	if !*previewMode {
		fmt.Println("Hello, Image Cut!!!")
		fmt.Println(getUUID())
		fmt.Println("Sleep 50 seconds...")
		time.Sleep(to * time.Second)
		os.Exit(0)
	}

	resizeImage(*width, *height, *imageFile, cfg.GetConfig())
}

func getUUID() uuid.UUID {
	return uuid.New()
}

func resizeImage(w, h int, fileName string, cfg models.Config) {
	p, err := preview.NewPreview(cfg.Previewer)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	img, ext, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	_ = file.Close()

	r := p.Preview(w, h, img, entities.ResizeOptions{
		Interpolation: cfg.Interpolation,
	})

	outFile := fmt.Sprintf("%s_resized_%dx%d.%s", fileName, w, h, ext)
	out, err := os.Create(outFile)

	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	switch ext {
	case "jpeg":
		err = jpeg.Encode(out, r, nil)
	case "png":
		err = png.Encode(out, r)
	default:
		fmt.Println("Unknown file type:", ext)
		os.Exit(fail)
	}

	if err != nil {
		log.Fatal("error encode image:", err)
	}

	fmt.Println("Resized successfully:", outFile)
}
