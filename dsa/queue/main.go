package main

import "fmt"

type Queue[T any] struct {
	element []T
}

func (q *Queue[T]) Push(ele T) {
	q.element = append(q.element, ele)
}

func (q *Queue[T]) Pop() (T, bool) {
	if len(q.element) == 0 {
		var zero T
		return zero, false
	}

	// 1. Pehle value nikaalo
	popped := q.element[0]

	// 2. Memory leak prevent karne ke liye 0th position zero-out karo
	var zero T
	q.element[0] = zero

	// 3. Slice aage badhao
	q.element = q.element[1:]

	return popped, true
}

func main() {
	var q Queue[int]

	q.Push(100)
	q.Push(200)

	val, ok := q.Pop()
	fmt.Println("Popped:", val, "Success:", ok) // 100, true

	val, ok = q.Pop()
	fmt.Println("Popped:", val, "Success:", ok) // 200, true
}
