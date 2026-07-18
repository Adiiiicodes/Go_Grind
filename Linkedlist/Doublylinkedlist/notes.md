# Doubly Linked List - DisplayBackward Notes

## Q: How does displaying backward work?

For displaying backward, the node with `next` value should be `nil`. Here's how we use this:

### First Loop (Finding the Tail)

```
for cur.next != nil {
    cur = cur.next
}
```

- Starts at the head
- Keeps moving forward while there IS a next node (`cur.next != nil`)
- Stops when `cur.next == nil` — meaning `cur` is now the last node (tail)

### Second Loop (Traversing Backward)

```
for cur != nil {
    print(cur.data, " ")
    cur = cur.prev
}
```

- Now we have the tail, traverse backward using `cur.prev`
- Stops when `cur == nil` (gone past the head)

### Example with list [10 → 20 → 30]

```
First loop finds:   10 → 20 → 30(next=nil) ✓ STOP HERE
                                   ↑
                                  cur

Second loop backward:
30 → 20 → 10 → nil
↑                ↑
cur starts     loop stops
```

We use the fact that the last node has `next == nil` to identify the tail, then leverage `prev` pointers to traverse backward.

---

## Q: Time Complexity - Can we optimize?

**Current complexity: O(n)** — and you can't do better when printing all n elements (must visit each node once).

### Optimization: Store a Tail Pointer

Instead of searching for the tail every time, maintain it in the struct:

```go
type DoublyLinkedList struct {
    head *Node
    tail *Node  // Add this
}
```

Update `Insert`:

```go
func (l *DoublyLinkedList) Insert(data int) {
    newNode := &Node{data: data}
    if l.head == nil {
        l.head = newNode
        l.tail = newNode  // Set tail
        return
    }
    cur := l.head
    for cur.next != nil {
        cur = cur.next
    }
    cur.next = newNode
    newNode.prev = cur
    l.tail = newNode  // Update tail
}
```

Optimized `DisplayBackward`:

```go
func (l *DoublyLinkedList) DisplayBackward() {
    cur := l.tail  // Start directly from tail
    for cur != nil {
        print(cur.data, " ")
        cur = cur.prev
    }
    println()
}
```

### Trade-off

| Aspect | Benefit |
|--------|---------|
| **Time** | Still O(n), but eliminates first traversal loop |
| **Space** | +8 bytes (one extra pointer) |
| **Access** | O(1) access to both head and tail |
