package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

// _____________________________________________________________________________________
type linkedList struct {
	head   *node
	length int
}

// _____________________________________________________________________________________
// adding prepend
func (a *linkedList) addPrependValue(val1 int) {
	nextNode := node{data: val1}
	// Header have value
	if a.head != nil {
		nextNode.next = a.head
		a.head = &nextNode
		a.length++
	}
}

// _____________________________________________________________________________________

// adding Value
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

// _____________________________________________________________________________________
// delet
func (l *linkedList) delet(val3 int) {
	if l.length == 0 {

		fmt.Println("No values in list")
	}

	target := l.head
	for target.next.data != val3 {
		target = target.next
	}
	target.next = target.next.next
	l.length--

}

// _____________________________________________________________________________________

// reverseing
func (s *linkedList) Reverse() {
	println("reversed orderd list :")
	curr := s.head
	var prev *node
	var temp *node

	for curr != nil {
		temp = curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}
	s.head = prev
}

// _____________________________________________________________________________________

// printing
func (l *linkedList) Print() {
	// if l.length == 0 {
	// 	fmt.Println("No values in list")
	// }
	print := l.head
	for i := 0; i < l.length; i++ {
		fmt.Print("  ", print.data)
		print = print.next
	}
	println("")
}

// _____________________________________________________________________________________
// _____________________________________________________________________________________
// _____________________________________________________________________________________
// _____________________________________________________________________________________
func main() {
	var input1, input2, input3, option, choice, count, size int
	var array [50]int
	myList := linkedList{}

	fmt.Println("")
	fmt.Println("default values in list are : 3, 4, 5, 6, 7")

	myList.insert(3)
	myList.insert(4)
	myList.insert(5)
	myList.insert(6)
	myList.insert(7)
	count = 1
	for i := 0; i < count; i++ {
		// user menu
		fmt.Println("choose an option")
		fmt.Println("1 : for adding a value")
		fmt.Println("2 : for add prepend")
		fmt.Println("3 : for deleting")
		fmt.Println("4 : for reversing")
		fmt.Println("5 : convert an array in to list")
		fmt.Scan(&choice)
		switch choice {

		case 1:
			// adding value
			fmt.Println("Enter a value")
			fmt.Scan(&input1)
			myList.insert(input1)
			fmt.Println("Current list is :")
			myList.Print()

		case 2:
			// adding prepend
			fmt.Println("Current list is :")
			myList.Print()
			fmt.Println("Enter a prepend value")
			fmt.Scan(&input2)
			myList.addPrependValue(input2)
			myList.Print()

		case 3:
			// deleting
			fmt.Println("Current list is :")
			myList.Print()
			fmt.Println("")
			fmt.Println("Enter a value to delet from node")
			fmt.Scanln(&input3)
			myList.delet(input3)
			myList.Print()

		case 4:
			// reversing
			fmt.Println("Current list is :")
			myList.Print()
			fmt.Println("")
			myList.Reverse()
			myList.Print()

		case 5:
			// arry to list
			fmt.Print("Enter the size of Array ")
			fmt.Scanln(&size)
			fmt.Println("Enter the values of array")
			for i := 0; i < size; i++ {
				fmt.Scan(&array[i])
				myList.insert(array[i])
			}
		default:
			fmt.Println("You entered wrong optionnumber.....")

			myList.Print()
		}
		fmt.Println(" if you want to continue enter 9 else 0")
		fmt.Scan(&option)
		if option == 9 {
			count++
		} else if option == 0 {
			println("")
			myList.Print()
			fmt.Println("program exit surcessfully")
		} else {
			fmt.Println("sorry you entered a wrong number .... program exit surcessfully")
			myList.Print()
		}
	}
	fmt.Println("thank you")

}
