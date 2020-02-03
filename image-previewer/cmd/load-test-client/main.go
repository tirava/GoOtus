/*
 * Project: Image Previewer
 * Created on 02.02.2020 16:20
 * Copyright (c) 2020 - Eugene Klimov
 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	fileName := filepath.Base(os.Args[0])
	flag.Usage = func() {
		fmt.Printf("Load test client for Preview image http server:\n"+
			"Parameters:\n"+
			"%s -urllist=urlfile -http=http://address:port/preview "+
			"-width=x -height=y -delay=x -cycles=x\n", fileName)
		fmt.Printf("Example:\n"+
			"%s -urllist=./examples/urlfile.txt "+
			"-http=http://localhost:8080/preview "+
			"-width=200 -height=100 -delay=100ms -cycles=10\n", fileName)
		flag.PrintDefaults()
	}

	paramURLFile := flag.String("urllist", "urlfile.txt", "path to file with source images URL list")
	paramHTTPServer := flag.String("http",
		"http://localhost:8080/preview", "scheme://address:port/path of the Previewer http-server")
	paramWidth := flag.Int("width", 100, "preview width")
	paramHeight := flag.Int("height", 100, "preview height")
	paramDelay := flag.String("delay", "0ms", "sleep between every url preview (in duration - s, ms, ns etc)")
	paramCycles := flag.Int("cycles", 1, "number of iterations all urls from list")
	flag.Parse()

	delay, err := parseDuration(*paramDelay)
	if err != nil {
		log.Fatal(err)
	}

	urlList, err := getURLsFromFile(*paramURLFile)
	if err != nil {
		log.Fatal(err)
	}

	startTest := time.Now()

	if err := startLoadTest(*paramHTTPServer, urlList, *paramWidth, *paramHeight, delay, *paramCycles); err != nil {
		log.Printf("Load test completed with errors:\n%s", err)
		// nolint
		os.Exit(1)
	}

	fmt.Printf("Load test completed successfully for time: %s\n", time.Since(startTest))
}

func parseDuration(s string) (time.Duration, error) {
	d, err := time.ParseDuration(s)
	if err != nil {
		return 0, err
	}

	return d, nil
}

func getURLsFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	urls := strings.Split(string(b), "\n")

	return urls, nil
}
