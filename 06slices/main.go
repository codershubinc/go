package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruit = []string{}
	fmt.Println("Type of fruitList is   ", fmt.Sprintf("%T", fruit))

	fruit = append(fruit, "Apple" , "Banana", "Grapes")
	fmt.Println("Fruit List after adding Apple: ", fruit)

	fruit = append(fruit[1:])
	fmt.Println("Fruit List after removing Apple: ", fruit)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 456
	highScores[2] = 678
	highScores[3] = 890
	fmt.Println("High Scores: ", highScores)
	
	highScores = append(highScores, 555, 666, 777)
	fmt.Println("High Scores after appending new scores: ", highScores)
	sort.Ints(highScores)
	fmt.Println("High Scores after sorting: ", highScores)

	

}
