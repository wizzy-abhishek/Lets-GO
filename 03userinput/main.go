package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	const welcome = "This is from user input"
	fmt.Println(welcome)

	// We use buffio package to take user input.

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter any string")

	//comma ok or err ok syntex
	input, _ /* we can write err, but we are not using it so use _*/ := reader.ReadString('\n')
	fmt.Println(input)

	// We have to convert the input to desired format to perform operations on it.
}
