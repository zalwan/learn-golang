package variables

import "fmt"

func VariablesAndConstants() {
	// Variable declaration
	// Manifest typing
	var firstName string = "Rizal" // Declaring a variable 'firstName' of type string
	var age int = 24               // Declaring a variable 'age' of type int

	// Short variable declaration
	// Type inference
	// This only works within functions
	city := "Tangerang"            // Declaring a variable 'city' with inferred type string
	city, distric := "Banten", 123 // Declaring 'city' and 'distric' with inferred types string and int

	// Declaration with keyword "new"
	withNew := new(bool) // Allocating memory for a boolean using the 'new' keyword

	// Constant declaration
	const religion = "Islam" // Declaring a constant 'religion' with the value "Islam"

	// Using the blank identifier for unused variable
	_ = "This is residu" // Special Variable: Blank Identifier (_)

	// Using 'make' to create a slice
	mySlice := make([]int, 3, 5) // Creating a slice of int with length 3 and capacity 5
	mySlice[0] = 1               // Assigning value to the first element of the slice
	mySlice[1] = 2               // Assigning value to the second element of the slice
	mySlice[2] = 3               // Assigning value to the third element of the slice

	// Using 'make' to create a map
	myMap := make(map[string]int) // Creating a map with string keys and int values
	myMap["Rizal"] = 24           // Adding key-value pair to the map
	myMap["Suryawan"] = 30        // Adding key-value pair to the map

	// Using 'make' to create a channel
	myChannel := make(chan int, 3) // Creating a channel of int with a buffer size of 2
	myChannel <- 42                // Sending a value to the channel
	myChannel <- 100               // Sending another value to the channel
	myChannel <- 10                // Sending another value to the channel

	_ = "This is residu" // Special Variable: Blank Identifier (_)

	// Output
	fmt.Println("name:", firstName)                    // Printing 'firstName'
	fmt.Println("age:", age)                           // Printing 'age'
	fmt.Println("city:", city+" "+fmt.Sprint(distric)) // Printing 'city' and 'district' with concatenation
	fmt.Println("religion:", religion)                 // Printing constant 'religion'
	fmt.Println("withNew:", *withNew)                  // Dereferencing 'withNew' pointer and printing the value

	// Output for 'make' examples
	fmt.Println("Slice:", mySlice)                  // Printing the slice
	fmt.Println("Map:", myMap)                      // Printing the map
	fmt.Println("Value from channel:", <-myChannel) // Receiving and printing the first value from the channel
	fmt.Println("Value from channel:", <-myChannel) // Receiving and printing the second value from the channel
}
