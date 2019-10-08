/*
 * HomeWork-6: FileCopy utility like dd
 * Created on 05.10.2019 14:26
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var fromFile, toFile string
var offset, limit int64

func init() {

	// set the custom Usage function
	fileName := filepath.Base(os.Args[0])
	flag.Usage = func() {
		fmt.Printf("usage: %s -from <source> -to <destination> [-offset bytes] [-limit bytes]\n", fileName)
		fmt.Printf("example: %s -from /path/to/source -to /path/to/dest -offset 1024 -limit 2048\n", fileName)
		flag.PrintDefaults()
	}

	// set flags
	flag.StringVar(&fromFile, "from", "", "file name to read from")
	flag.StringVar(&toFile, "to", "", "file name to write to")
	flag.Int64Var(&offset, "offset", 0, "offset in input file, bytes")
	flag.Int64Var(&limit, "limit", 0, "limit, bytes")
}

func main() {
	flag.Parse()

	// no blank path
	if fromFile == "" || toFile == "" {
		flag.Usage()
		os.Exit(2)
	}

	n, err := CopyFileSeekLimit(os.Stdout, toFile, fromFile, offset, limit)
	if err != nil {
		log.Fatalln("error copy data:", err)
	}
	fmt.Printf("\nCopied %d bytes from offset %d\n", n, offset)
}

// CopyFileSeekLimit copies limit bytes from position offset in src file to dst file
// and returns successfully copied bytes and errors
// w is writer for progress
func CopyFileSeekLimit(w io.Writer, dst, src string, offset, limit int64) (int64, error) {

	from, err := os.Open(src)
	if err != nil {
		return 0, fmt.Errorf("can't open source file: %s", err)
	}
	defer from.Close()

	to, err := os.Create(dst)
	if err != nil {
		return 0, fmt.Errorf("can't create destination file: %s", err)
	}
	defer to.Close()

	if _, err := from.Seek(offset, io.SeekStart); err != nil {
		return 0, fmt.Errorf("can't set seeker position: %s", err)
	}

	// set limit = file size if it = 0
	if limit == 0 {
		stat, err := from.Stat()
		if err != nil {
			return 0, fmt.Errorf("can't get file stat: %s", err)
		}
		limit = stat.Size()
	}

	// set buffer 100 chunks for simple percentage
	bufSize := limit / 100
	if bufSize == 0 {
		bufSize = 1
	}

	// set start position
	lr := io.LimitReader(from, limit)
	buf := make([]byte, bufSize)
	var count int64

	for {
		n, err := lr.Read(buf)
		if err != nil && err != io.EOF {
			return 0, fmt.Errorf("can't read from file: %s", err)
		}
		if n == 0 {
			break
		}
		if _, err := to.Write(buf[:n]); err != nil {
			return 0, fmt.Errorf("can't write to file: %s", err)
		}
		count += int64(n)
		// uncomment for small file for visible progress
		//time.Sleep(time.Millisecond * 100)
		percent := count * 100 / limit
		if _, err := fmt.Fprintf(w, "Copied: %d%%\r", percent); err != nil {
			log.Println("can't write progress to writer")
		}
	}

	return count, nil
}
