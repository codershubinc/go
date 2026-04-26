package main

import "fmt"

func main(){

	fmt.Println("go with go")

	age := 30;
	fmt.Println("age is", age)

	var name string = "John"
	fmt.Println("name is", name)

	var height float64 = 5.9
	fmt.Println("height is", height)

	isEmployed := true
	fmt.Println("is employed:", isEmployed)

	// Using a constant
	const pi = 3.14
	fmt.Println("value of pi is", pi)

	// Using a rune
	var char rune = 'A'
	fmt.Println("character is", string(char))

	// Using a byte
	var b byte = 65 // ASCII value of 'A'
	fmt.Println("byte value is", b)
}