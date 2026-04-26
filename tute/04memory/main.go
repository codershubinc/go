package main

import "fmt"

func main() {
	fmt.Println("Its a pointers")
	var ptr *int
	fmt.Println("Value of ptr is: ", ptr)

	num := 23
	ptr = &num
	fmt.Println("Value of ptr is: ", ptr)
	fmt.Println("Value at address ptr is: ", *ptr)
	
	*ptr = *ptr + 2
	fmt.Println("Value at address ptr is: ", *ptr)
	fmt.Println("Value of num is: ", num)

}
