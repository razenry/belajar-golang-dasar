package main

import "fmt"

func main() {
	// Hello World
	fmt.Println("Hello, World!")

	// Variables
	var name string = "Golang"
	var age int = 10
	isAwesome := true
	fmt.Printf("Name: %s, Age: %d, Is Awesome: %t\n", name, age, isAwesome)

	// Constants
	const pi = 3.14
	fmt.Println("Value of Pi:", pi)

	// Arrays
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", numbers)

	// Slices
	slice := []int{10, 20, 30}
	slice = append(slice, 40)
	fmt.Println("Slice:", slice)

	// Map
	person := map[string]string{
		"name": "John",
		"age":  "30",
	}
	fmt.Println("Map:", person)

	// If-Else
	if age > 18 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Minor")
	}

	// For Loop
	for i := 0; i < 5; i++ {
		fmt.Println("Loop:", i)
	}

	// Functions
	result := add(5, 3)
	fmt.Println("Addition Result:", result)

	// Struct
	type User struct {
		Name  string
		Email string
	}
	user := User{Name: "Alice", Email: "alice@example.com"}
	fmt.Println("User Struct:", user)
}

// Function to add two numbers
func add(a int, b int) int {
	return a + b
}