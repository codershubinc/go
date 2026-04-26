// Basic Concept: Pointers and Mutability
//
// Task: You are writing a user profile updater for your backend.
// The `UpdateEmail` function is supposed to change the user's email,
// but right now, it has a bug. The original email remains unchanged!
//
// 1. Look at `UpdateEmail`. Why isn't it modifying the original struct?
// 2. Fix the `UpdateEmail` function signature so that it can modify the original `User`.
// 3. Update the `main` function so it correctly passes `u` to your fixed function.

package main

import "fmt"

type User struct {
	Username string
	Email    string
}

// BUG: This function doesn't actually change the user's email in main!
func UpdateEmail(u *User, newEmail string) {
	u.Email = newEmail
	fmt.Println("Inside function, email is:", u.Email)
}

func main() {
	u := User{
		Username: "swapnil",
		Email:    "old@codershubinc.com",
	}

	fmt.Println("Before update:", u.Email)

	// Try to update the email
	UpdateEmail(&u, "new@codershubinc.com")

	fmt.Println("After update:", u.Email)
	// EXPECTED: new@codershubinc.com
	// ACTUAL: old@codershubinc.com
}
