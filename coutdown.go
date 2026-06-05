package main

import (
	"fmt"
	"time"
)

func coutdown(duration time.Duration, msg string) {
	for duration >= 0 {
		// \r go to the start of the line
		// \033[K erase everythong on the lin

		fmt.Print("\r\033[K")
		fmt.Print("\r")
		fmt.Printf("%s%d Seconds", msg, int(duration.Seconds()))

		time.Sleep(time.Second)
		duration -= time.Second
	}
}
