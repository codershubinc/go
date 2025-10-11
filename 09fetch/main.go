package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("Fetching ......")

	// fetch api
	res, err := http.Get("https://openapi.codershubinc.tech/v1.0/user")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Parse JSON response
	var apiResponse APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
		return
	}

	// Display parsed data
	fmt.Printf("Status Code: %d\n", apiResponse.StatusCode)
	fmt.Printf("Success: %t\n", apiResponse.Success)
	fmt.Printf("Message: %s\n", apiResponse.Message)

	user := apiResponse.Data.User
	fmt.Printf("\n--- User Information ---\n")
	fmt.Printf("ID: %s\n", user.ID)
	fmt.Printf("Full Name: %s\n", user.Name.FullName)
	fmt.Printf("First Name: %s\n", user.Name.FirstName)
	fmt.Printf("Middle Name: %s\n", user.Name.MiddleName)
	fmt.Printf("Last Name: %s\n", user.Name.LastName)
	fmt.Printf("Title: %s %s\n", user.Name.Prefix, user.Name.Title)
	fmt.Printf("Gender: %s\n", user.Name.Gender)
	fmt.Printf("Country Code: %s\n", user.Name.CountryCode)

	fmt.Printf("\n--- Preferences ---\n")
	fmt.Printf("Mode: %s\n", user.Prefs.Mode)
	fmt.Printf("Browser: %s\n", user.Prefs.Devices.Browser)
	fmt.Printf("OS: %s\n", user.Prefs.Devices.OS)
	fmt.Printf("Location: %s\n", user.Prefs.Devices.Location)
	fmt.Printf("Hobbies: %v\n", user.Prefs.Info.Hobby)

	address := apiResponse.Data.Address
	fmt.Printf("\n--- Address ---\n")
	fmt.Printf("Street: %d %s\n", address.Street.Number, address.Street.Name)
	fmt.Printf("City: %s\n", address.City)
	fmt.Printf("State: %s\n", address.State)
	fmt.Printf("Post Code: %d\n", address.PostCode)
	fmt.Printf("Country: %s (%s)\n", address.Country.Name.Country, address.Country.Code)
	fmt.Printf("Coordinates: %s, %s\n", address.Coordinates.Latitude, address.Coordinates.Longitude)
	fmt.Printf("Time Zone: %s\n", address.TimeZone.Name)
	fmt.Printf("Current Time: %s\n", address.TimeZone.Zone)

	fmt.Printf("\n--- Document ID ---\n")
	fmt.Printf("Document ID: %s\n", apiResponse.Data.DocumentID)

	fmt.Println("\nDone")
}
