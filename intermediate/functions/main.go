package main

import "fmt"

// 1. Basic function
func add(a, b int) int {
	return a + b
}

// 2. Multiple return values
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b

	return quotient, remainder
}

// 3. Named return values
func calculate(a, b int) (sum int, difference int) {
	sum = a + b
	difference = a - b

	return
}

// 4. Variadic function
func sum(numbers ...int) int {
	total := 0

	for _, n := range numbers {
		total += n
	}

	return total
}

// 5. Function as parameter
func calculateOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// 6. Function returning a function / Closure
func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

// 7. Function working with pointer
func updateValue(x *int) {
	*x = 200
}

type Server struct {
	Name string
	Port int
}

// 8. Function working with struct pointer
func updatePort(s *Server) {
	s.Port = 9090
}

func main() {

	// -----------------------------
	// 1. Basic function
	// -----------------------------

	result := add(10, 20)

	fmt.Println("Add:", result)

	// -----------------------------
	// 2. Multiple return values
	// -----------------------------

	quotient, remainder := divide(10, 3)

	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)

	// Ignore one return value using _
	quotient, _ = divide(20, 3)

	fmt.Println("Only quotient:", quotient)

	// -----------------------------
	// 3. Named return values
	// -----------------------------

	sumResult, difference := calculate(20, 10)

	fmt.Println("Sum:", sumResult)
	fmt.Println("Difference:", difference)

	// -----------------------------
	// 4. Variadic function
	// -----------------------------

	fmt.Println("Sum:", sum(10, 20))
	fmt.Println("Sum:", sum(10, 20, 30, 40))

	// Passing a slice to variadic function
	numbers := []int{10, 20, 30}

	fmt.Println("Slice sum:", sum(numbers...))

	// -----------------------------
	// 5. Function as a value
	// -----------------------------

	multiply := func(a, b int) int {
		return a * b
	}

	fmt.Println("Multiply:", multiply(10, 5))

	// -----------------------------
	// 6. Function as an argument
	// -----------------------------

	addFunction := func(a, b int) int {
		return a + b
	}

	result = calculateOperation(10, 20, addFunction)

	fmt.Println("Operation result:", result)

	// We can also pass function directly
	result = calculateOperation(10, 20, func(a, b int) int {
		return a * b
	})

	fmt.Println("Direct function result:", result)

	// -----------------------------
	// 7. Closure
	// -----------------------------

	counter1 := counter()

	fmt.Println("Counter1:", counter1())
	fmt.Println("Counter1:", counter1())
	fmt.Println("Counter1:", counter1())

	// Another closure has its own state
	counter2 := counter()

	fmt.Println("Counter2:", counter2())
	fmt.Println("Counter2:", counter2())

	// -----------------------------
	// 8. Function + pointer
	// -----------------------------

	x := 100

	updateValue(&x)

	fmt.Println("Updated value:", x)

	// -----------------------------
	// 9. Function + struct pointer
	// -----------------------------

	server := Server{
		Name: "web",
		Port: 8080,
	}

	updatePort(&server)

	fmt.Println("Server port:", server.Port)
}
