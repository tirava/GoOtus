/*
 * HomeWork-1: Hello now()
 * Created on 04.09.19 19:28
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
)

var serversNTP = [...]string{
	"ntp5.stratum2.ru",
	"time-nw.nist.gov",
	"ru.pool.ntp.org",
}

func main() {

	errCode := 0

	for _, server := range serversNTP {
		now, err := ntp.Time(server)
		if err != nil {
			log.Println("Server:", server, "get time error:", err)
			errCode = 1
			continue
		}
		errCode = 0
		fmt.Println("Server:", server, "Time:", now)
	}

	os.Exit(errCode)
}
