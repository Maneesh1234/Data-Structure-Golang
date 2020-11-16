package main

import "fmt"

func push(stack []int, element int) []int {
	stack = append(stack, element) // Simply append to push.
	fmt.Println("pushed:", element)
	return stack
}

func pop(stack []int) []int {
	element := stack[len(stack)-1] // The first element is the one to be popd.
	fmt.Println("poped:", element)
	return stack[:len(stack)-1] // Slice off the element once it is popd.
}

func main() {
	var stack []int // Make a stack of ints.

	stack = push(stack, 10)
	stack = push(stack, 20)
	stack = push(stack, 30)

	fmt.Println("Stack:", stack)

	stack = pop(stack)
	fmt.Println("stack:", stack)

	stack = push(stack, 40)
	fmt.Println("stack:", stack)
}
