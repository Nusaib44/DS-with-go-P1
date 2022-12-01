package main

import "fmt"

type stack struct {
	data   []int
	length int
}

// push
func (a *stack) push() {
	var stack_limit, val1 int

	fmt.Println("enter the limit of stack")
	fmt.Scan(&stack_limit)
	fmt.Println("enter value")
	if a.length == 0 {
		fmt.Scan(&val1)
		a.data = append(a.data, val1)
	}
	for i := 0; i < stack_limit-1; i++ {
		fmt.Scan(&val1)
		a.data = append(a.data, val1)
	}

}

// pop
func (a *stack) pop() int {

	last := len(a.data) - 1
	removed := a.data[last]
	a.data = a.data[:last]
	// fmt.Println(removed)
	return removed
}

// print
func (s *stack) print() {
	v := s.data
	fmt.Println(v)

}

func main() {

	mystack := stack{}
	fmt.Println("Hello welcome....")
	mystack.push()
	mystack.print()
	// mystack.pop()
	fmt.Println(mystack.pop())
	mystack.print()

}
