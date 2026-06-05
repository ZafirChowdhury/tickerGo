package main

import (
	"fmt"
	"time"
)

func main() {
	duration := 10 * time.Second

	for duration >= 0 {
		// \r go to the start of the line
		// \033[K erase everythong on the line
		fmt.Print("\r\033[K")

		fmt.Printf("\rTime remaning: %d seconds", int(duration.Seconds()))

		time.Sleep(time.Second)
		duration -= time.Second
	}

	fmt.Println("\nDone")
}
