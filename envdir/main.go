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

	// no blank path
	if execFile == "" || envDir == "" {
		flag.Usage()
		os.Exit(2)
	}

	err := EnvDirExec(execFile, envDir)
	if err != nil {
		log.Fatalln("error execution:", err)
	}
	//fmt.Printf("\nCopied %d bytes from offset %d\n", n, offset)
}
