package datastructure

// Queue type definition
type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func (q *Queue) isEmpty() bool {
	return q.length == 0
}

// Length method return the length of the queue
func (q *Queue) Length() int {
	return q.length
}

// EnQueue an element to the tail of the queue
func (q *Queue) EnQueue(str string) {
	n := &Node{str, nil}
	// method 1
	// 	if q.tail == nil {
	// 		q.tail = n
	// 	} else {
	// 		q.tail.Next = n
	// 		if q.head == nil {
	// 			q.head = q.tail
	// 		}
	// 		q.tail = n
	// 	}
	// method 2
	// 	if q.length == 0 {
	// // 		n.Next = q.tail // redundant
	// 		q.head = n
	// 		q.length++
	// 		return
	// 	}
	// 	if q.length == 1 {
	// 		q.tail = n
	// 		q.head.Next = n
	// 		q.length++
	// 		return
	// 	}
	// 	q.tail.Next = n
	// 	q.tail = n
	// 	q.length++
	// 	return
	// method 3
	switch q.length {
	case 0:
		q.head = n
	case 1:
		q.tail = n
		q.head.Next = n
	default:
		q.tail.Next = n
		q.tail = n
	}
	q.length++
	return
}

// DeQueue an element from the head of the queue
func (q *Queue) DeQueue() (string, bool) {
	// 	if q.isEmpty() {
	// 		return "", false
	// 	}
	// 	if q.head == nil && q.tail != nil {
	// 		str := q.tail.Value
	// 		q.tail = nil
	// 		return str, true
	// 	}
	// 	if q.head.Next == q.tail {
	// 		str := q.head.Value
	// 		q.head = nil
	// 		return str, true
	// 	}
	// 	n := q.head.Next
	// 	str := q.head.Value
	// 	q.head = n
	// 	return str, true
	if q.length == 0 {
		return "", false
	}
	str := q.head.Value
	newHead := q.head.Next
	q.head = newHead
	q.length--
	return str, true
}

// String function define custom output
func (q Queue) String() string {
	if q.length == 0 {
		return ""
	}
	n := q.head
	str := n.Value
	for n.Next != nil {
		n = n.Next
		str += (" - " + n.Value)
	}
	// 	n := q.head
	// 	str := n.Value
	// 	for q.length > 0 { // here is a potential bug, if q.length == 1, n.Next is nil
	// 		n = n.Next
	// 		str += (" - " + n.Value)
	// 	}
	return str
}
