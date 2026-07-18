package main

type Node struct {
	data int
	next *Node
	
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(data int) {
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
}

func (l *LinkedList) Display() {
	cur := l.head
	for cur != nil {
		print(cur.data, " ")
		cur = cur.next
	}
	println()
}

func main() {
	list := &LinkedList{}
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Display() // Output: 10 20 30
}
