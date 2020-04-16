package main

import (
	"log"

	http "gitlab.com/tirava/shop/internal/http_server"
)

func main() {
	//fmt.Println("Waiting sleep...")
	//time.Sleep(10 * time.Second) //  nolint
	serv := http.NewServer(":8000")
	if err := serv.StartServer(); err != nil {
		log.Fatal(err)
	}
}
