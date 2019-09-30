/*
 * HomeWork-5: Worker Pool
 * Created on 28.09.19 22:11
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package workerpool implements N-workers with stopping after X-errors.
// Commented lines is for benchmark results, uncomment if no need benchmarks.
package workerpool

import (
	"errors"

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

	// check errors from jobs
	go func() {
		countErr := 0
		for range errChan {
			countErr++
			// fmt.Printf("\tTotal number of errors - %d, MAX errors: %d\n", countErr, maxErrors)
			if countErr >= maxErrors {
				// fmt.Printf("\tTotal number of errors: %d, MAX errors: %d, aborting all jobs ...\n", countErr, maxErrors)
				close(abortChan) // abort all workers
				return
			}
		}
	}()

	// start workers
	for i := 0; i < maxJobs; i++ {
		// i := i
		eg.Go(func() error {
			for job := range jobsChan {
				select {
				case <-abortChan:
					// fmt.Printf("\tWorker '%d' aborted\n", i)
					return ErrWorkerAborted
				default:
					// fmt.Printf("\tWorker '%d' started\n", i)
					if err := job(); err != nil {
						// fmt.Println(err)
						errChan <- err
					}
					// fmt.Printf("\tWorker '%d' finished\n", i)
				}
			}
			// fmt.Printf("\tWorker '%d' exited\n", i)
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
