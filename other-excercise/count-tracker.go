package main

import "fmt"

// Declare variable activeUserCount
// your code goes here
var activeUserCount = 0

func entry() {
	activeUserCount++
	// Hint: you can use the "++" operator to increment a variable by 1
	// your code goes here
}

func exit() {
	activeUserCount--
	// Hint: you can use the "--" operator to decrement a variable by 1
	// your code goes here
}

func main() {
	entry()
	entry()
	exit()
	exit()
	entry()
	entry()
	fmt.Println(activeUserCount)
}
