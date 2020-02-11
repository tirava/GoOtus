package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"

	"gitlab.com/tirava/image-previewer/internal/configs"
	"gitlab.com/tirava/image-previewer/internal/domain/entities"
	"gitlab.com/tirava/image-previewer/internal/domain/preview"
	"gitlab.com/tirava/image-previewer/internal/helpers"
)

const (
	previewConfigPath = "PREVIEWER_CONFIG_PATH"
)

func main() {
	fileName := filepath.Base(os.Args[0])
	flag.Usage = func() {
		fmt.Printf("Preview image tool w/o starting http server:\n"+
			"%s [-config=configFile|inmemory] -width=x -height=y -file=imageFile\n"+
			"[PREVIEWER_CONFIG_PATH=configFile|inmemory] %s -width=x -height=y -file=imageFile\n",
			fileName, fileName)
		flag.PrintDefaults()
	}

	config := flag.String("config", "config.yml", "path to yaml config file or 'inmemory'")
	imageFile := flag.String("file", "", "path to image file")
	width := flag.Int("width", 0, "preview width")
	height := flag.Int("height", 0, "preview height")
	flag.Parse()

	if *width == 0 || *height == 0 || *imageFile == "" {
		flag.Usage()
		// nolint:gomnd
		os.Exit(2)
	}

	if os.Getenv(previewConfigPath) != "" {
		*config = os.Getenv(previewConfigPath)
	}

	cfg, err := configs.NewConfig(*config)
	if err != nil {
		log.Fatal(err)
	}

	conf := cfg.GetConfig()
	conf.Cacher = "nolimit"
	conf.Storager = "inmemory"
	conf.StoragePath = ""
	conf.MaxCacheItems = 0

	prev, err := helpers.InitPreview(conf)
	if err != nil {
		log.Fatal(err)
	}

	opts := entities.ResizeOptions{
		Interpolation: conf.Interpolation,
	}

	if *config == "inmemory" {
		fmt.Println("InMemory config:")
		fmt.Println("Previewer:", conf.Previewer)
		fmt.Println("Interpolation:", "NearestNeighbor")
	}

	if err := resizeImage(*width, *height, *imageFile, *prev, opts); err != nil {
		log.Fatal(err)
	}
}

func resizeImage(w, h int, fileName string, prev preview.Preview, opts entities.ResizeOptions) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("unable to open image file: %w", err)
	}

	img, ext, err := image.Decode(file)
	if err != nil {
		return err
	}

	_ = file.Close()
	r := prev.Preview(w, h, img, opts)

	outFile := fmt.Sprintf("%s_resized_%dx%d.%s", fileName, w, h, ext)
	out, err := os.Create(outFile)

	if err != nil {
		return err
	}
	defer out.Close()

	switch ext {
	case "jpeg":
		err = jpeg.Encode(out, r, nil)
	case "png":
		err = png.Encode(out, r)
	case "gif":
		err = gif.Encode(out, r, nil)
	default:
		return fmt.Errorf("unknown file type: %s", ext)
	}

	if err != nil {
		return fmt.Errorf("error encode image: %w", err)
	}

	fmt.Println("Resized successfully:", outFile)

	return nil
}
