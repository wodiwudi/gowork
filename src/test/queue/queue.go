package queue

import "errors"

type Queue []interface{}

func (q *Queue) Push(i interface{}) {
	*q = append(*q, i)
}
func (q *Queue) Pop() (interface{}, error) {
	if q.isEmpty() == true {
		return 0, errors.New("queue is empty")
	}
	tail := (*q)[0]
	*q = (*q)[1:len(*q)]
	return tail, nil
}
func (q *Queue) isEmpty() bool {
	if len(*q) == 0 {
		return true
	}
	return false
}
