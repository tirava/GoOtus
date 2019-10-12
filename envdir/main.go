/*
 * HomeWork-7: envdir utility like envdir
 * Created on 11.10.2019 21:50
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var (
	execFile, envDir string
	inheritEnv       bool
)

func init() {

	// set the custom Usage function
	fileName := filepath.Base(os.Args[0])
	flag.Usage = func() {
		fmt.Printf("usage: %s -env <envdir> -exec <filename> -inherit\n", fileName)
		fmt.Printf("example: %s -env /path/to/dir -exec /path/to/file\n", fileName)
		flag.PrintDefaults()
	}

	// set flags
	flag.StringVar(&execFile, "exec", "", "file name to execution")
	flag.StringVar(&envDir, "env", "", "directory where env vars are")
	flag.BoolVar(&inheritEnv, "inherit", false, "inherit system env variables")
}

func main() {
	flag.Parse()

	// run as env client if no params
	if execFile == "" && envDir == "" {
		for _, v := range os.Environ() {
			fmt.Println(v)
		}
		os.Exit(0)
	} else if execFile == "" || envDir == "" { // no blank path if params exist
		flag.Usage()
		os.Exit(2)
	}

	err := EnvDirExec(os.Stdout, envDir, execFile, inheritEnv)
	if err != nil {
		log.Fatalln(err)
	}
}

// How to test:
// go build .
// go test -v

// How to use:
// ./envdir -env /full/path/to/dir -exec /full/path/to/envdir [-inherit]
// ./envdir -env /full/path/to/dir -exec env [-inherit]
