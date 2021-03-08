package queue

// Queue data structure
type Queue struct {
	Queue []interface{}
}

// New is conustrctor of Queue
func (q *Queue) New() {
	q.Queue = make([]interface{}, 0)
}

func (q *Queue) enqueue(item interface{}) {
	tempQueue := []interface{}{item}
	q.Queue = append(tempQueue, q.Queue...)
}

func (q *Queue) dequeue() interface{} {
	length := len(q.Queue)
	item := q.Queue[length-1]
	q.Queue = q.Queue[:length-1]
	return item
}

func (q *Queue) isEmpty() bool {
	return len(q.Queue) == 0
}

func (q *Queue) size() int {
	return len(q.Queue)
}
