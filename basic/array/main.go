package main

import "fmt"

func main() {
	numbers := [5]int{10, 20, 30, 40, 50}
	fmt.Println(numbers[0])
	fmt.Println(len(numbers)) // 5

	largest := -100
	for _, val := range numbers {
		if val > largest {
			largest = val
		}
	}
	fmt.Println(largest)

	arr := []int{1, 2, 3}
	arr = append(arr, 4)
	arr = append(arr, 5)

	fmt.Println(cap(arr))

	arr = append(arr, 12)
	arr = append(arr, 12)
	arr = append(arr, 12)
	arr = append(arr, 12)
	arr = append(arr, 12)

	fmt.Println(cap(arr))

}
