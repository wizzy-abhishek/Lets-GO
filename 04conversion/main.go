package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome")
	fmt.Println("Please rate between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	var input, _ = reader.ReadString('\n')
	numericalValue, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(numericalValue + 1)
	}
}
