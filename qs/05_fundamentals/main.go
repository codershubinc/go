// Basic Concept: Maps and the "Comma-OK" Idiom
//
// Task: You are writing an authentication helper for your backend.
//
// 1. Complete `GetRole`. Use the "comma-ok" idiom to check if `username` is in the `db` map.
// 2. If the user exists, return their role and `nil` for the error.
// 3. If the user does NOT exist, return an empty string `""` and a new error saying "user not found".
//    (Hint: Use `errors.New("...")` to create an error).

package main

import (
	"errors"
	"fmt"
)

func GetRole(db map[string]string, username string) (string, error) {
	role, ok := db[username]
	if !ok {
		return "", errors.New("User not found")
	}
	return role, nil
}

func main() {
	// Our mock database of users and their roles
	db := map[string]string{
		"alice": "admin",
		"bob":   "", // Bob exists in the DB, but has no role assigned yet
	}

	// Test 1: Alice (Should return "admin", nil)
	role, err := GetRole(db, "alice")
	fmt.Printf("Alice   -> Role: '%s', Error: %v\n", role, err)

	// Test 2: Bob (Should return "", nil)
	role, err = GetRole(db, "bob")
	fmt.Printf("Bob     -> Role: '%s', Error: %v\n", role, err)

	// Test 3: Charlie (Should return "", "user not found")
	role, err = GetRole(db, "charlie")
	fmt.Printf("Charlie -> Role: '%s', Error: %v\n", role, err)
}
