/*
 * HomeWork-4: Worker Pool
 * Created on 26.09.19 22:11
 * Copyright (c) 2019 - Eugene Klimov
 */
// Package workerpool implements N-workers with stopping after X-errors.
//package workerpool
package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const (
	JOBTIMEOUT = 8   // in seconds
	JOBSNUM    = 100 // number of all jobs
	MAXJOBS    = 10  // max concurrency jobs/workers
	MAXERRORS  = 15  // max errors from all jobs
)

type Job func() error

// WorkerPool is the main worker pool manager.
func WorkerPool(jobs []Job, maxJobs int, maxErrors int) error {
	var eg errgroup.Group

	jobsChan := make(chan Job, maxJobs)
	errChan := make(chan error, maxJobs)
	abortChan := make(chan bool)

	// check errors from jobs
	go func() {
		countErr := 0
		for range errChan {
			countErr++
			fmt.Printf("\tTotal number of errors - %d, MAX errors: %d\n", countErr, maxErrors)
			if countErr >= maxErrors {
				fmt.Printf("\tTotal number of errors: %d, MAX errors: %d, aborting all jobs ...\n", countErr, maxErrors)
				close(abortChan) // abort all workers
				return
			}
		}
	}()

	// start workers
	for i := 0; i < maxJobs; i++ {
		i := i
		eg.Go(func() error {
			for job := range jobsChan {
				select {
				case <-abortChan:
					fmt.Printf("\tWorker '%d' aborted\n", i)
					//return err // what scenario for error?
					return nil
				default:
					break
				}
				fmt.Printf("\tWorker '%d' started\n", i)
				if err := job(); err != nil {
					fmt.Println(err)
					errChan <- err
				}
				fmt.Printf("\tWorker '%d' finished\n", i)
			}
			fmt.Printf("\tWorker '%d' exited\n", i)
			//return err // what scenario for error?
			return nil
		})
	}

	// send jobs to workers
	for _, j := range jobs {
		select {
		case <-abortChan:
			break
		default:
			jobsChan <- j
		}
	}

	close(jobsChan)
	err := eg.Wait()
	close(errChan)

	return err
}

func main() {
	rand.Seed(time.Now().UnixNano())
	jobs := make([]Job, 0)

	// jobs slice
	for i := 0; i < JOBSNUM; i++ {
		i := i
		job := func() error {
			d := rand.Intn(3) + 1                      // random time for every job
			n := strconv.Itoa(i)                       // job id
			time.Sleep(time.Duration(d) * time.Second) // any work here
			if rand.Intn(2) == 0 {                     // error gen randomly
				return fmt.Errorf("job '%s' returned error", n)
			}
			fmt.Printf("job '%s' ended successfully, duration: %d seconds\n", n, d)
			return nil
		}
		jobs = append(jobs, job)
	}

	// start
	if err := WorkerPool(jobs, MAXJOBS, MAXERRORS); err != nil {
		log.Fatalln("One or more jobs returned with errors!")
	}
	fmt.Println("All jobs returned successfully!")
}
