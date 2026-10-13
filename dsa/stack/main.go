package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	arr  []int
	size int
	top  int
}

func (st *Stack) getSize() int {
	return st.top + 1
}

func (st *Stack) push(ele int) {
	st.top += 1
	st.arr = append(st.arr, ele)
}

func (st *Stack) pop() (int, error) {
	if st.top == -1 {
		return -1, errors.New("Stack is Empty")
	}
	ele := st.arr[st.top]
	st.top--
	lastIndex := len(st.arr) - 1
	st.arr = st.arr[:lastIndex]
	return ele, nil

}

func main() {
	st := Stack{
		size: 0,
		top:  -1,
	}

	fmt.Println("Stack Size is ", st.getSize())
	ele, err := st.pop()
	if err == nil {
		fmt.Println(err)
	} else {
		fmt.Println(ele)
	}
	st.push(10)
	st.push(20)
	st.push(30)
	st.push(40)

	fmt.Println("Stack Size is ", st.getSize())
	ele, err = st.pop()
	if err == nil {
		fmt.Println("Pop Element", ele)
	} else {
		fmt.Println(err)
	}

}
