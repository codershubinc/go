package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Grapes"
	fruitList[3] = "Orange"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Length of Fruit List is: ", len(fruitList))
	
	var vegList = [3]string{"Potato", "Tomato", "Beans"}
	fmt.Println("Veg List is: ", vegList)
	fmt.Println("Length of Veg List is: ", len(vegList))	

	

}
