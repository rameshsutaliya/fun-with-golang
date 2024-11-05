package main

import (
	"fmt"
	"fun-with-golang/ch1/functions"
)

func main() {
	functions.Greet()                  // it should print the "Hello World!"
	fmt.Println(functions.Named(7))    // should split into two values as 3, and 4
	fmt.Println(functions.Add(13, 17)) // should return 30
	fmt.Println(functions.Swap(1, 2))  // should print 2, 1
}
