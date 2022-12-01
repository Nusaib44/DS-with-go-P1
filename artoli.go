package main

import "fmt"

type node struct {
	data int
	next *node
}

// _____________________________________________________________________________________
type linkedList struct {
	head   *node
	length int
}

func (a *linkedList) insert(val2 int) {

	nextNode := node{data: val2}
	add := a.head
	// head is nil
	if a.length == 0 {
		a.head = &nextNode
		a.length++
		return
	}
	// head have value
	// addig value to next node
	for i := 0; i < a.length; i++ {
		if add.next == nil {
			add.next = &nextNode
			a.length++
			return
		}
		add = add.next
	}
}

func (l *linkedList) Print() {
	// if l.length == 0 {
	// 	fmt.Println("No values in list")
	// }
	print := l.head
	fmt.Println("Values in linked list are :")
	for i := 0; i < l.length; i++ {
		fmt.Print("  ", print.data)
		print = print.next
	}
	println("")
}

func main() {
	myList := linkedList{}
	var array [50]int
	var size int
	fmt.Print("Enter the size of Array ")
	fmt.Scanln(&size)
	fmt.Println("Enter the values of array")
	for i := 0; i < size; i++ {
		fmt.Scan(&array[i])
		myList.insert(array[i])
	}

	myList.Print()
}
