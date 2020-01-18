/*
 * Project: Image Previewer
 * Created on 10.01.2020 13:20
 * Copyright (c) 2020 - Eugene Klimov
 */

package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"time"

	"gitlab.com/tirava/image-previewer/internal/domain/preview"

	"github.com/google/uuid"
)

const to = 50

func main() {
	fmt.Println("Hello, Image Cut!!!")
	fmt.Println(getUUID())

	// nolint
	//resizeJPG(1024, 252)
	//resizeJPG(333, 666)
	//resizeJPG(256, 126)
	//resizeJPG(500, 500)
	resizeJPG(2000, 1000)

	fmt.Println("Sleep 50 seconds...")
	time.Sleep(to * time.Second)
}

func getUUID() uuid.UUID {
	return uuid.New()
}

func resizeJPG(w, h int) {
	p, err := preview.NewPreview("nfnt_crop")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open("examples/_gopher_original_1024x504.jpg")
	if err != nil {
		log.Fatal(err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	_ = file.Close()

	r := p.Preview(w, h, img)

	out, err := os.Create("examples/test_resized.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	if err := jpeg.Encode(out, r, nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Resized successfully.")
}
