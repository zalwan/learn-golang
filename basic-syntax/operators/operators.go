package operators

import "fmt"

func Operators() {
	// Integer variables for arithmetic operations
	a, b := 10, 5

	// Floating-point variables for arithmetic operations
	x, y := 7.5, 2.5

	// --- Arithmetic Operators ---
	// Addition, Subtraction, Multiplication, Division, Modulus
	fmt.Println("Arithmetic Operators:")
	fmt.Println("a + b =", a+b) // 10 + 5 = 15
	fmt.Println("a - b =", a-b) // 10 - 5 = 5
	fmt.Println("a * b =", a*b) // 10 * 5 = 50
	fmt.Println("a / b =", a/b) // 10 / 5 = 2
	fmt.Println("a % b =", a%b) // 10 % 5 = 0

	// Using floating-point values
	fmt.Println("x + y =", x+y) // 7.5 + 2.5 = 10.0
	fmt.Println("x - y =", x-y) // 7.5 - 2.5 = 5.0
	fmt.Println("x * y =", x*y) // 7.5 * 2.5 = 18.75
	fmt.Println("x / y =", x/y) // 7.5 / 2.5 = 3.0

	// --- Comparison Operators ---
	// Equal to, Not equal to, Greater than, Less than, Greater than or equal to, Less than or equal to
	fmt.Println("\nComparison Operators:")
	fmt.Println("a == b:", a == b) // false
	fmt.Println("a != b:", a != b) // true
	fmt.Println("a > b:", a > b)   // true
	fmt.Println("a < b:", a < b)   // false
	fmt.Println("a >= b:", a >= b) // true
	fmt.Println("a <= b:", a <= b) // false

	// --- Logical Operators ---
	// AND (&&), OR (||), NOT (!)
	isActive, isVerified := true, false
	fmt.Println("\nLogical Operators:")
	fmt.Println("isActive && isVerified:", isActive && isVerified) // false
	fmt.Println("isActive || isVerified:", isActive || isVerified) // true
	fmt.Println("!isActive:", !isActive)                           // false
	fmt.Println("!isVerified:", !isVerified)                       // true
}
