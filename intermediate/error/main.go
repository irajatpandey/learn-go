package main

import (
	"errors"
	"fmt"
)

var ErrUserNotFound = errors.New("user not found")

// errors.New()
// fmt.Errorf()
// Returning error
func getUser(id int) (string, error) {

	if id <= 0 {
		return "", errors.New("invalid user ID")
	}

	if id == 101 {
		return "", fmt.Errorf("user with ID %d not found", id)
	}

	return "Rajat", nil
}

// Error wrapping with %w
func getUserData() error {

	err := ErrUserNotFound

	return fmt.Errorf("failed to get user data: %w", err)
}

func main() {

	// --------------------------------
	// 1. errors.New()
	// --------------------------------

	err1 := errors.New("something went wrong")

	fmt.Println("Error 1:", err1)

	// --------------------------------
	// 2. fmt.Errorf()
	// --------------------------------

	id := 101

	err2 := fmt.Errorf("user %d not found", id)

	fmt.Println("Error 2:", err2)

	// --------------------------------
	// 3. Function returning error
	// --------------------------------

	name, err := getUser(101)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("User:", name)
	}

	// --------------------------------
	// 4. %w - Error wrapping
	// --------------------------------

	err3 := getUserData()

	fmt.Println("Error 3:", err3)

	// --------------------------------
	// 5. errors.Is()
	// --------------------------------

	if errors.Is(err3, ErrUserNotFound) {
		fmt.Println("User not found error detected")
	}
}
