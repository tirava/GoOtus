/*
 * HomeWork-5: Worker Pool
 * Created on 29.09.19 11:22
 * Copyright (c) 2019 - Eugene Klimov
 */

package workerpool

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var testCases = []struct {
	jobs        []Job
	jobsNum     int // number of all jobs
	maxJobs     int // max concurrency jobs/workers
	maxJobsTime int // in milliseconds - random jobs run time from 0 to maxJobsTime
	maxErrors   int // max errors from all jobs
	errExpected error
	description string
}{
	{
		jobsNum:     20,
		maxJobs:     5,
		maxJobsTime: 100,
		maxErrors:   1,
		errExpected: ErrWorkerAborted,
		description: "20 jobs, 5 workers, max 1 errors, workers return errors",
	},
	{
		jobsNum:     10,
		maxJobs:     3,
		maxJobsTime: 1000,
		maxErrors:   9,
		errExpected: nil,
		description: "10 jobs, 3 workers, max 9 errors, no workers errors",
	},
	{
		jobsNum:     100,
		maxJobs:     10,
		maxJobsTime: 10,
		maxErrors:   55,
		errExpected: nil,
		description: "100 jobs, 10 workers, max 50 errors, more loading",
	},
}

func TestWorkerPool(t *testing.T) {
	genJobs()
	for _, test := range testCases {
		err := WorkerPool(test.jobs, test.maxJobs, test.maxErrors)
		if err != test.errExpected {
			if err != nil {
				t.Errorf("FAIL '%s':\n\t WorkerPool returned error '%s', expected nil error.", test.description, err)
			} else {
				t.Errorf("FAIL '%s':\n\t WorkerPool returned nil error, expected error '%s'.", test.description, ErrWorkerAborted)
			}
			continue
		}
		t.Logf("PASS WorkerPool - '%s'", test.description)
	}
}

func genJobs() {
	rand.Seed(time.Now().UnixNano())

	for i, test := range testCases {
		test.jobs = make([]Job, 0)
		for i := 0; i < test.jobsNum; i++ {
			i := i
			t := test.maxJobsTime
			job := func() error {
				d := rand.Intn(t) + 1                           // random time for every job
				n := strconv.Itoa(i)                            // job id
				time.Sleep(time.Duration(d) * time.Millisecond) // any work here
				if rand.Intn(2) == 0 {                          // error gen randomly
					return fmt.Errorf("job '%s' returned error", n)
				}
				//fmt.Printf("job '%s' ended successfully, duration: %d ms\n", n, d)
				return nil
			}
			test.jobs = append(test.jobs, job)
		}
		testCases[i].jobs = test.jobs
	}
}

func BenchmarkWorkerPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			_ = WorkerPool(test.jobs, test.maxJobs, test.maxErrors)
		}
	}
}
