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

var execFile, envDir string

func init() {

	// set the custom Usage function
	fileName := filepath.Base(os.Args[0])
	flag.Usage = func() {
		fmt.Printf("usage: %s -exec <filename> -env <envdir>\n", fileName)
		fmt.Printf("example: %s -exec /path/to/file -env /path/to/dir\n", fileName)
		flag.PrintDefaults()
	}

	// set flags
	flag.StringVar(&execFile, "exec", "", "file name to execution")
	flag.StringVar(&envDir, "env", "", "directory where env vars are")
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

	err := EnvDirExec(execFile, envDir)
	if err != nil {
		log.Fatalln("error execution:", err)
	}
}

// ./envdir -exec /full/path/to/envdir -env /full/path/to/dir
