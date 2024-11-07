package functions

import "fmt"

/**
* function: Set of instruction set.
* Function may take zero or more inputs, and may return zero or more outputs
 */

// Greet the world.
// Greet starts with a capital letter, so it is exportable.
func Greet() {
	fmt.Println("Hello World!")
}

// Add adds two integer numbers
func Add(num1 int, num2 int) int {
	return num1 + num2
}

// Swap multiple returns
func Swap(num1 int, num2 int) (int, int) {
	// this can also be return as: Swap(num1, num2 int) (int, int)
	return num2, num1
}

// Named returns the value in two halves
/*
Go's return values may be named. If so, they are treated as variables defined at the top of the function
*/
func Named(val int) (firstPart, secondPart int) {
	firstPart = val / 2
	secondPart = val - firstPart
	return
}

func GetGreetMsg() string {
	return "˗ˏˋ ★ ˎˊ˗\n welcome !!"
}
