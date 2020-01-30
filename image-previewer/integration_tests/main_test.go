/*
 * Project: Image Previewer
 * Created on 30.01.2020 16:20
 * Copyright (c) 2020 - Eugene Klimov
 */

package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/godog"
)

func TestMain(m *testing.M) {
	fmt.Println("Wait 5s for services availability...")
	// nolint
	time.Sleep(5 * time.Second)

	status := godog.RunWithOptions("integration", func(s *godog.Suite) {
		FeatureContext(s)
	}, godog.Options{
		Format:    "pretty",
		Paths:     []string{"features"},
		Randomize: 0,
	})

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
