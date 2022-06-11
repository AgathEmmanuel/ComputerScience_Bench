package main


import (
	"fmt"
//	"github.com/golang-collections/collections/queue"
)


type (
	Queue struct {
		start, end *node
		length int
	}
	node struct {
		value interface{}
		next *node
	}
)


// Take the next item off the front of the queue
func (this *Queue) Dequeue() interface{} {
	if this.length == 0 {
		return nil
	}
	n := this.start
	if this.length == 1 {
		this.start = nil
		this.end = nil
	} else {
		this.start = this.start.next
	}
	this.length--
	return n.value
}
// Put an item on the end of a queue
func (this *Queue) Enqueue(value interface{}) {
	n := &node{value,nil}
	if this.length == 0 {
		this.start = n
		this.end = n		
	} else {
		this.end.next = n
		this.end = n
	}
	this.length++
}
// Return the number of items in the queue
func (this *Queue) Len() int {
	return this.length
}
// Return the first item in the queue without removing it
func (this *Queue) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.start.value
}

func main() {
	fmt.Println("hello")


	q:=Queue{nil,nil,0}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue("one")
	q.Enqueue(3)
	q.Enqueue(3)
	q.Enqueue(3)
	fmt.Println(q.Len())
	fmt.Println(q.Peek())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}



