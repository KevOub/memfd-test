package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0
	fmt.Print("Hello Dave!")
	for {
		time.Sleep(time.Second * 5)
		fmt.Printf("COUNTER: %d\n", counter)
	}
}
