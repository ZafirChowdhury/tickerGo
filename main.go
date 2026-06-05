package main

import (
	"fmt"
	"time"
)

func main() {
	duration := 10 * time.Second
	coutdown(duration, "Time remaning: ")
	fmt.Println("\nDone")
}
