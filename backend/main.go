package main

import (
	"fmt"
	"os"
)

func main() {
	var APIkey string = ""
	fmt.Println("Hello, welcome to the Go application!")
	fmt.Println("to continue, please follow the instructions provided.")
	fmt.Println("please provide a apikey.")
	fmt.Scanln(&APIkey)
	fmt.Printf("You have entered: %s\n", APIkey)
	fmt.Println("Thank you for providing the API key.")
	menu()

}
func menu() {
	var selector int
	fmt.Println("please select your option ")
	fmt.Println("1 = enqueue prompt")
	fmt.Println("2 = Show History")
	fmt.Println("3 = View Responses")
	fmt.Println("4 = Exit")
	fmt.Scanln(&selector)
	switch selector {
	case 1:
		promptenqueue()
	case 2:
		History()
	case 3:
		ViewResponse()
	case 4:
		os.Exit(0)
	}
}

func promptenqueue() {
	var prompt string
	fmt.Println("please give your prompt")
	fmt.Scanln(&prompt)
	fmt.Println("you have entered %s!\n", prompt)

	var selector int
	fmt.Println("please select your option ")
	fmt.Println("1 = return to menu")
	fmt.Println("2 = exit")
	fmt.Scanln(&selector)
	switch selector {
	case 1:
		menu()
	case 2:
		os.Exit(0)
	}

}
func History() {

	fmt.Println("fetched history:")

	var selector int
	fmt.Println("please select your option ")
	fmt.Println("1 = return to menu")
	fmt.Println("2 = exit")
	fmt.Scanln(&selector)
	switch selector {
	case 1:
		menu()
	case 2:
		os.Exit(0)
	}
}
func ViewResponse() {
	fmt.Println("fetching responses")

	var selector int
	fmt.Println("please select your option ")
	fmt.Println("1 = return to menu")
	fmt.Println("2 = exit")
	fmt.Scanln(&selector)
	switch selector {
	case 1:
		menu()
	case 2:
		os.Exit(0)
	}

}

// This is a simple Go program that prints "Hello, World!" to the console.
//test
// It serves as a basic example of a Go application structure.
//asdasdasd
