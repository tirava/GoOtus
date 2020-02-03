/*
 * Project: Image Previewer
 * Created on 02.02.2020 16:56
 * Copyright (c) 2020 - Eugene Klimov
 */

package main

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"
)

func startLoadTest(httpServer string, urlList []string, width, height int,
	delay time.Duration, cycles int) error {
	for i := 1; i <= cycles; i++ {
		startCycle := time.Now()

		var eg errgroup.Group

		for _, url := range urlList {
			// nolint
			if len(url) < 3 { // no fake strings
				continue
			}

			url := url

			time.Sleep(delay)

			eg.Go(func() error {
				fullPath := path.Join(httpServer, strconv.Itoa(width), strconv.Itoa(height), url)
				fullPath = strings.ReplaceAll(fullPath, ":/", "://")
				// nolint
				resp, err := http.Get(fullPath)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				return nil
			})
		}

		if err := eg.Wait(); err != nil {
			return err
		}

		fmt.Printf("Iteration '#%d' completed for time: %s\n", i, time.Since(startCycle))
	}

	return nil
}
