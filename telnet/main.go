/*
 * HomeWork-10: telnet client
 * Created on 05.11.2019 22:03
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func main() {

	timeoutArg := flag.String("timeout", "10s", "timeout for connection (duration)")
	fileName := filepath.Base(os.Args[0])
	flag.Usage = func() {
		fmt.Printf("usage: %s [--timeout] <host> <port>\n", fileName)
		fmt.Printf("example1: %s 1.2.3.4 567\n", fileName)
		fmt.Printf("example2: %s --timeout=10s 8.9.10.11 1213\n", fileName)
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) < 2 {
		flag.Usage()
		os.Exit(2)
	}
	timeout, err := time.ParseDuration(*timeoutArg)
	if err != nil {
		log.Fatalln(err)
	}
	addr := flag.Arg(0) + ":" + flag.Arg(1)

	ctx, cancel := context.WithCancel(context.Background())

	dialer := &net.Dialer{Timeout: timeout}
	conn, err := dialer.Dial("tcp", addr)
	if err != nil {
		log.Fatalln("Cannot connect:", err)
	}
	fmt.Println("Connected to:", addr)
	fmt.Println("Press 'Ctrl+D or Ctrl+C' for exit")

	abort := make(chan bool)
	go readRoutine(ctx, conn, abort)
	ch := make(chan string)
	go writeRoutine(ctx, conn, ch, abort)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		sig := <-c
		fmt.Println("Got signal:", sig)
		abort <- true
	}()

	<-abort
	cancel()

	time.Sleep(1 * time.Second) // wait 0.5 second for every socket goroutine
	close(ch)

	fmt.Println("Closing connection... ")
	if err := conn.Close(); err != nil {
		log.Fatalln("Error close connection:", err)
	}
	fmt.Println("...closed connection")
	fmt.Println("Exited.")
}

func readRoutine(ctx context.Context, conn net.Conn, abort chan bool) {
	reply := make([]byte, 1)
OUTER:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Exiting from reading...")
			break OUTER
		default:
			// set deadline for read socket - need for 'select loop' continue
			if err := conn.SetReadDeadline(time.Now().Add(500 * time.Millisecond)); err != nil {
				log.Println(err)
			}
			n, err := conn.Read(reply)
			if err != nil {
				if err == io.EOF {
					fmt.Println("Remote host aborted connection, exiting from reading...")
					abort <- true
					break OUTER
				}
				if netErr, ok := err.(net.Error); ok && !netErr.Timeout() {
					log.Println(err)
				}
			}
			if n == 0 {
				break
			}
			fmt.Print(string(reply))
		}
	}
	fmt.Println("...exited from reading")
}

func writeRoutine(ctx context.Context, conn net.Conn, ch chan string, abort chan bool) {
	go func(ch chan<- string) {
		reader := bufio.NewReader(os.Stdin)
		for {
			s, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					fmt.Println("Ctrl+D detected, aborting...")
					abort <- true
					return
				}
				log.Println(err)
			}
			ch <- s
		}
	}(ch)

OUTER:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Exiting from writing...")
			break OUTER
		default:

		STDIN:
			for {
				select {
				case stdin, ok := <-ch:
					if !ok {
						break STDIN
					}
					if _, err := conn.Write([]byte(stdin)); err != nil {
						log.Println(err)
					}
					// wait deadline for input
				case <-time.After(500 * time.Millisecond):
					break STDIN
				}
			}
		}
	}
	fmt.Println("...exited from writing")
}

//func main() {
//
//	fmt.Println("Launching server...")
//
//	// listen on all interfaces
//	ln, _ := net.Listen("tcp", ":8081")
//
//	// accept connection on port
//	conn, _ := ln.Accept()
//
//	// run loop forever (or until ctrl-c)
//	for {
//		// will listen for message to process ending in newline (\n)
//		message, _ := bufio.NewReader(conn).ReadString('\n')
//		// output message received
//		fmt.Print("Message Received:", string(message))
//		// sample process for string received
//		newmessage := strings.ToUpper(message)
//		// send new string back to client
//		conn.Write([]byte(newmessage + "\n"))
//	}
//}
