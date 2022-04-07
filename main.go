package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var table = make(chan string)

func main() {
	wg.Add(2)
	go ping()
	go pong()
	wg.Wait()
}

func ping() {
	for {
		table <- "Ping..."
		var pong = <-table
		fmt.Printf("%v\n", pong)
		time.Sleep(1 * time.Second)
	}
}

func pong() {
	for {
		var ping = <-table
		fmt.Printf("%v\n", ping)
		time.Sleep(1 * time.Second)
		table <- "Pong..."
	}
}
