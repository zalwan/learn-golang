package datatypes

import "fmt"

func DataTypes() {
	// Integer types
	var age int = 25                   // 'int' is a signed integer type. Size depends on the system (32-bit or 64-bit).
	var year int32 = 2024              // 'int32' is a 32-bit signed integer.
	var largeNumber int64 = 1234567890 // 'int64' is a 64-bit signed integer.

	// Unsigned integer types
	var positiveAge uint = 30         // 'uint' is an unsigned integer type (only positive numbers).
	var bigNumber uint64 = 9876543210 // 'uint64' is an unsigned 64-bit integer.

	// Floating point types
	var pi float32 = 3.14     // 'float32' is a 32-bit floating-point number.
	var gravity float64 = 9.8 // 'float64' is a 64-bit floating-point number (higher precision).

	// Complex types
	var complexNumber complex64 = 3 + 4i // 'complex64' is a 64-bit complex number.
	var bigComplex complex128 = 2 + 7i   // 'complex128' is a 128-bit complex number.

	// String type
	var firstName string = "Rizal"    // 'string' is a sequence of characters.
	var message string = "Hello, Go!" // Another string.

	// Boolean type
	var isActive bool = true // 'bool' can either be true or false.

	// Array type
	var numbers [3]int = [3]int{1, 2, 3} // 'array' type with 3 integer elements.

	// Slice type
	var fruitList []string = []string{"Apple", "Banana", "Cherry"} // 'slice' type, dynamic array.

	// Map type
	var userDetails map[string]string = map[string]string{
		"name": "Rizal",
		"age":  "24",
	} // 'map' type is a key-value pair collection.

	// Pointer type
	var ptr *int = &age // 'pointer' type holds the memory address of the variable.

	// Constant type
	const piConstant float64 = 3.14159 // 'const' to define constant values.

	// Output values of different data types
	fmt.Println("Integer:", age)
	fmt.Println("Year:", year)
	fmt.Println("Large Number:", largeNumber)
	fmt.Println("Positive Age:", positiveAge)
	fmt.Println("Big Number:", bigNumber)
	fmt.Println("Pi (float32):", pi)
	fmt.Println("Gravity (float64):", gravity)
	fmt.Println("Complex Number:", complexNumber)
	fmt.Println("Big Complex:", bigComplex)
	fmt.Println("First Name:", firstName)
	fmt.Println("Message:", message)
	fmt.Println("Is Active:", isActive)
	fmt.Println("Numbers Array:", numbers)
	fmt.Println("Fruit List Slice:", fruitList)
	fmt.Println("User Details Map:", userDetails)
	fmt.Println("Pointer to Age:", ptr)
	fmt.Println("Pi Constant:", piConstant)
}
