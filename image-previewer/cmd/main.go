/*
 * Project: Image Previewer
 * Created on 10.01.2020 13:20
 * Copyright (c) 2020 - Eugene Klimov
 */

package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const to = 50

func main() {
	fmt.Println("Hello, Image Cut!!!")
	fmt.Println(getUUID())
	fmt.Println("Sleep 50 seconds...")
	time.Sleep(to * time.Second)
}

func getUUID() uuid.UUID {
	return uuid.New()
}
