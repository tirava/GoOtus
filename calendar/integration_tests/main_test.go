/*
 * HomeWork-9: Integration tests
 * Created on 13.12.2019 13:08
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	fmt.Println("Wait 10s for services availability...")
	time.Sleep(10 * time.Second)

	status := godog.RunWithOptions("integration", func(s *godog.Suite) {
		godog.SuiteContext(s)
		FeatureContextAddEvent(s)
		FeatureContextListEvents(s)
		FeatureContextQueueEvent(s)
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
