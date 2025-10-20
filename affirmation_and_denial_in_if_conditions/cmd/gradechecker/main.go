package main

import (
	"flag"
	"fmt"

	"example.com/gradechecker/internal/auth"
	"example.com/gradechecker/internal/grade"
)

func main() {
	// Define command-line flags
	name := flag.String("name", "Guest", "User name")
	isAdmin := flag.Bool("admin", false, "Is the user an admin")
	isLoggedIn := flag.Bool("login", false, "Is the user logged in")

	flag.Parse()

	// Create user based on flags
	user := auth.NewUser(*name, *isAdmin, *isLoggedIn)

	// Display greeting
	fmt.Println(user.GetGreeting())

	// Check authentication
	fmt.Println(user.GetStatusMessage())

	if !user.CanAccessGrades() {
		fmt.Println("You must be logged in to check grades.")
		return
	}

	// Get grade input
	var score int
	fmt.Print("Please enter your grade to check if you passed or not: ")
	fmt.Scanln(&score)

	// Display grade result
	fmt.Println(grade.DisplayResult(score))

	// Show admin-only features if applicable
	if user.CanAccessAdminFeatures() {
		fmt.Println("\nADMIN INFORMATION:")
		fmt.Println(grade.AdminGradeStats(score))
		fmt.Println("You can access additional admin features here.")
	}
}
