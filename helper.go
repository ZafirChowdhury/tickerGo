package main

import (
	"fmt"
)

func printUsage(invalid bool) {
	if invalid {
		fmt.Println("Invalid usage")
	}

	fmt.Println("Only use int numbers to respond")
}

func takeInput(msg string) (int, error) {
	var input int

	fmt.Println(msg)
	fmt.Print("> ")

	_, err := fmt.Scan(&input)
	if err != nil {
		return 0, err
	}

	return input, err
}
