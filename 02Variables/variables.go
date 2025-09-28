package main

import "fmt"

var token = "12233u1" //allowed
var token2 int = 1234

const LoggedToken string = "1084nkjwdn" //The capital first letter denotes it is public

func main() {
	var myVar string //empty string is the default value ""
	fmt.Println(myVar)
	fmt.Printf("The variable is of type %T \n", myVar)

	var isLogged bool
	fmt.Println(isLogged)
	fmt.Printf("The variable isLogged is of type %T \n", isLogged)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("The variable isLogged is of type %T \n", smallInt)

	var smallFloat float32 = 255.455555555
	fmt.Println(smallFloat)
	fmt.Printf("The variable isLogged is of type %T \n", smallFloat)

	// default values and aliases.

	// for bool is false.
	// for int it is 0.

	// IMPLICIT TYPE

	var website = 12 // Lexer decides the type, but it is strongly typed we can not change the type after this.
	fmt.Println(website)

	// no var style using walarus operator :=
	// walarus operator is allowed inside a method only.

	randomValue := 1020
	fmt.Println(randomValue)

	fmt.Println(LoggedToken)
	fmt.Printf("%T \n", LoggedToken)

}
