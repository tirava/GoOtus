package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/DATA-DOG/godog"
)

const ready = 25 * time.Second

func TestMain(m *testing.M) {
	fmt.Printf("\nWait %s for services availability...\n", ready)
	fmt.Println("Sleep is needed for readiness network interfaces " +
		"(for Linux it must be 5+s but MacOS X requires 15+s)")
	time.Sleep(ready)

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
