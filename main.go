package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	printUsage(false)

	work, err := takeInput("How much do you want to work per sessions? (In minutes)")
	if err != nil {
		log.Println(err.Error())
		return
	}

	rest, err := takeInput("How long do you want your breaks? (In minutes)")
	if err != nil {
		log.Println(err.Error())
		return
	}

	sessions, err := takeInput("How many sessions do you want to do?")
	if err != nil {
		log.Println(err.Error())
		return
	}

	workFlag := true
	count := 0
	for count < sessions {
		fmt.Printf("Sessons remaning: %d\n", (sessions - count))
		if workFlag {
			coutdown(time.Duration(work)*time.Minute, "Time until your work seasson is done: ")
			fmt.Print("\r\033[K")
			fmt.Println("Go take a break!")
		}

		coutdown(time.Duration(rest)*time.Minute, "Time until your break ends: ")
		fmt.Print("\r\033[K")
		fmt.Println("Get back to work!")

		count++
		workFlag = false
	}

	fmt.Println("\nCongratulation on getting your work done")
}
