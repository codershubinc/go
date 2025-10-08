package main

import "fmt"

func main() {

	var fruit = []string{}
	fmt.Println("Type of fruitList is   ", fmt.Sprintf("%T", fruit))

	fruit = append(fruit, "Apple" , "Banana", "Grapes")
	fmt.Println("Fruit List after adding Apple: ", fruit)

	fruit = append(fruit[1:])
	fmt.Println("Fruit List after removing Apple: ", fruit)

}
