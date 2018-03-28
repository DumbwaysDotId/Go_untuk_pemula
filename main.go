package main

import (
	"fmt"
)

func printAllVariables() {
	var helloWorld string = "Hello World Go"
	var fruitCount int = 2
	var fruitPrice float64 = 2500.20
	var isOldFruit bool = false

	fmt.Printf("%v\n%v\n%v\n%v\n", helloWorld, fruitCount, fruitPrice, isOldFruit)
}

func helloWorld() string {
	return "Hello World Go"
}

func add(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}

func main() {
	// printAllVariables()
	// fmt.Println(helloWorld())
	fmt.Println(add(3, 4))
}
