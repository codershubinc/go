package main

import "fmt"

// defining a struct
type User struct {
	Name  string
	Email string
	Age   int
}

func main() {

	// creating an instance of the struct
	user1 := User{
		Name:  "Alice",
		Email: "alice@example.com",
		Age:   30,
	}

	fmt.Println("User 1:", user1.Name)
	fmt.Printf(" Details: %+v\n", user1)

}
