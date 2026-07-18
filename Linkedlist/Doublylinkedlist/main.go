package main

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
}

func (l *DoublyLinkedList) Insert(data int) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
		return
	}
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newNode
	newNode.prev = cur
}

func (l *DoublyLinkedList) Display() {
	cur := l.head
	for cur != nil {
		print(cur.data, " ")
		cur = cur.next
	}
	println()
}

func (l *DoublyLinkedList) DisplayBackward() {
	if l.head == nil {
		return
	}
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	for cur != nil {
		print(cur.data, " ")
		cur = cur.prev
	}
	println()
}

func main() {
	list := &DoublyLinkedList{}
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	print("Forward: ")
	list.Display()
	print("Backward: ")
	list.DisplayBackward()
}
