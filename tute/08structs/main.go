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

	// input struct values at runtime
	var user2 User
	fmt.Println("Enter user details:")
	fmt.Print("Name: ")
	fmt.Scan(&user2.Name)
	fmt.Print("Email: ")
	fmt.Scan(&user2.Email)
	fmt.Print("Age: ")
	fmt.Scan(&user2.Age)

	fmt.Println("User 2:", user2.Name)
	fmt.Printf(" Details: %+v\n", user2)

}
