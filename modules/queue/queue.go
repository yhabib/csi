// Package queue implements the Queue ADS with two Stacks
package queue

import (
	"github.com/yhabib/csi/tree/main/modules/stack"
)

// Queue data structure
type Queue struct {
	rearStack  stack.Stack
	frontStack stack.Stack
}

// New is conustrctor of Queue
func (q *Queue) New() {
	q.items = make([]interface{}, 0)
}

func (q *Queue) enqueue(item interface{}) {
	tempQueue := []interface{}{item}
	q.items = append(tempQueue, q.items...)
}

func (q *Queue) dequeue() interface{} {
	length := len(q.items)
	item := q.items[length-1]
	q.items = q.items[:length-1]
	return item
}

func (q *Queue) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) size() int {
	return len(q.items)
}
