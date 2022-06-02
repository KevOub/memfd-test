package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// get os pid
	pid := os.Getpid()
	fmt.Printf("[*] PID: %d\n", pid)

	// counter := 0
	fmt.Print("Hello Dave!")
	for {
		time.Sleep(time.Second * 5)
		// fmt.Printf("COUNTER: %d\n", counter)
		// counter += 1
	}
}
