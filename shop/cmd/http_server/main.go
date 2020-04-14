package main

import (
	"log"

	http "gitlab.com/tirava/shop/internal/http_server"
)

func main() {
	serv := http.NewServer(":8000")
	if err := serv.StartServer(); err != nil {
		log.Fatal(err)
	}
}
