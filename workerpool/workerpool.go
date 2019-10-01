/*
 * HomeWork-5: Worker Pool
 * Created on 28.09.19 22:11
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package workerpool implements N-workers with stopping after X-errors.
package workerpool

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
)

// Job is type for jobs
type Job func() error

// ErrWorkerAborted is an error to return on an aborted workers.
var ErrWorkerAborted = errors.New("workers aborted")

// WorkerPool is the main worker pool manager.
func WorkerPool(jobs []Job, maxJobs int, maxErrors int) error {
	var eg errgroup.Group

	jobsChan := make(chan Job, maxJobs)
	errChan := make(chan error, maxJobs)
	abortChan := make(chan bool)
	msgChan := make(chan string, maxJobs)
	defer close(errChan)
	defer close(msgChan)

	// check messages from workers
	go func() {
		for msg := range msgChan {
			fmt.Print(msg)
		}
	}()

	// check errors from workers jobs
	go func() {
		countErr := 0
		for err := range errChan {
			fmt.Println(err)
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
					msgChan <- fmt.Sprintf("\tWorker '%d' aborted\n", i)
					return ErrWorkerAborted
				default:
					msgChan <- fmt.Sprintf("\tWorker '%d' started\n", i)
					if err := job(); err != nil {
						errChan <- err
					}
					msgChan <- fmt.Sprintf("\tWorker '%d' finished\n", i)
				}
			}
			msgChan <- fmt.Sprintf("\tWorker '%d' exited\n", i)
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
	close(jobsChan) // time to return

	return eg.Wait()
}
