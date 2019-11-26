package main

import "fmt"

func main() {
	// Define a map
	emails := make(map[string]string)

	// Assign key/values
	emails["Chef"] = "chef@email.com"
	emails["Sharon"] = "sharon@email.com"
	emails["Randy"] = "randy@email.com"

	fmt.Println(emails)
	fmt.Println(len(emails))

	// Delete from map
	delete(emails, "Chef")
	fmt.Println(emails)

	// Declare and declare together
	otherEmails := map[string]string{"Stan": "stan@email.com", "Strong": "strong@email.com"}
	fmt.Println(otherEmails)
}
