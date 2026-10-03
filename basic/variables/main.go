package main

import "fmt"

// Variables In the Go Lang
func main() {

	const PI = 3.1415
	fmt.Println(PI)

	const (
		MONDAY    = 1
		TUESDAY   = 2
		WEDNESDAY = 3
	)

	fmt.Println(MONDAY)

	name := "Rajat"
	age := 31
	salary := 20.5
	isDevOps := true

	fmt.Println(name, " ")
	fmt.Println(age)
	fmt.Println(salary)
	fmt.Println(isDevOps)

	var a int = 10
	var b float64 = 20.5

	var c float64 = float64(a) + b
	fmt.Println(c)

	//var age1 int = 30
	//var salary2 float64 = float64(age1)
}
