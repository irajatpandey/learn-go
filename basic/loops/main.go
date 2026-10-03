package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While Loop
	i := 0

	for i < 5 {
		fmt.Println(i)
		i++
	}
	for i := 0; i < 11; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	numbers := []int{10, 20, 30}

	for index, value := range numbers {
		fmt.Println(index, value)
	}
	for _, value := range numbers {
		fmt.Println(value)
	}

	for index, value := range numbers {
		fmt.Println(index, value)
	}
}
