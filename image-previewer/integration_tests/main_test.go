package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/godog"
)

func TestMain(m *testing.M) {
	// sleep is need for ready network interfaces (for Linux it to 5s but MacOS requires 10+s)
	fmt.Println("Wait 10s for services availability...")
	// nolint:gomnd
	time.Sleep(10 * time.Second)

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
