// Basic Concept: Interfaces and Polymorphism
//
// Task: You are building an alerting system for your backend.
// You want to be able to send alerts via Email or SMS, without
// hardcoding the specific service into your core alerting logic.
//
// 1. Define an interface named `Alerter` that requires a single method: `Alert(message string)`
// 2. Add an `Alert(message string)` method to the `EmailService` struct.
// 3. Add an `Alert(message string)` method to the `SMSService` struct.
// 4. In `main`, call the `SendUrgentAlert` function using both services.

package main

import "fmt"

// YOUR CODE HERE: 1. Define the Alerter interface
type Alerter interface {
	Alert(message string)
}

type EmailService struct {
	EmailAddress string
}

// YOUR CODE HERE: 2. Implement the Alert method for EmailService
// It should print something like: "Sending Email to [address]: [message]"
func (e EmailService) Alert(message string) {
	fmt.Println("Sending Email to", e.EmailAddress, ":", message)
}

type SMSService struct {
	PhoneNumber string
}

// YOUR CODE HERE: 3. Implement the Alert method for SMSService
// It should print something like: "Sending SMS to [phone]: [message]"
func (s SMSService) Alert(message string) {
	fmt.Println("Sending SMS to", s.PhoneNumber, ":", message)
}

// This function accepts ANY struct that implements the Alerter interface!
// You do not need to change this function.
func SendUrgentAlert(a Alerter, msg string) {
	fmt.Println("[SYSTEM ALARM]")
	a.Alert(msg)
}

func main() {
	email := EmailService{EmailAddress: "admin@codershubinc.com"}
	sms := SMSService{PhoneNumber: "555-0199"}

	// YOUR CODE HERE: 4. Call SendUrgentAlert passing the `email` service and a message
	SendUrgentAlert(email, "email notification")

	// YOUR CODE HERE: 5. Call SendUrgentAlert passing the `sms` service and a message
	SendUrgentAlert(sms, "Sms msg")

}
