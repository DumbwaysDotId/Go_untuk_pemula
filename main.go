package main

import (
	"fmt"
)

// var helloWorld string = "Hello World Go outside scope"
// var helloWorld = "Hello World Go outside scope"

func main() {
	helloWorld := "Hello World Go"
	fruitCount := 5
	fmt.Println("%v %v", helloWorld, fruitCount)
}
