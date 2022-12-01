package main

import "fmt"

type node struct {
	data string
	prev *node
	next *node
}

type doublyLinkedList struct {
	len  int
	tail *node
	head *node
}

func (d *doublyLinkedList) AddFrontNode(data string) {
	newNode := &node{data: data}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.len++
	return
}

func (d *doublyLinkedList) AddEndNode(data string) {
	newNode := &node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		d.tail = newNode
	}
	d.len++
	return
}

func (d *doublyLinkedList) TraverseForward() {
	temp := d.head
	for temp != nil {
		fmt.Printf("value = %v \n", temp.data)
		temp = temp.next
	}
	fmt.Println()
}

func (d *doublyLinkedList) TraverseReverse() {
	temp := d.tail
	for temp != nil {
		fmt.Printf("value = %v", temp.data)
		temp = temp.prev
		fmt.Println()
	}

}

func (d *doublyLinkedList) Size() int {
	return d.len
}
func main() {
	doublyList := doublyLinkedList{}

	doublyList.AddFrontNode("a")
	doublyList.AddEndNode("D")
	doublyList.AddFrontNode("b")
	doublyList.AddFrontNode("c")

	doublyList.AddEndNode("E")

	fmt.Printf("Size of doubly linked ist: %d\n", doublyList.Size())

	doublyList.TraverseForward()
	doublyList.TraverseReverse()
}
