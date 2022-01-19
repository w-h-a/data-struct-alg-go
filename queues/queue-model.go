package queues

import (
	"errors"
)

var (
	ErrQueueIsNil = errors.New("queue is nil")
)

type Queuer interface {
	Enq(x int) *Queue
	Head() (int, error)
	Deq() (*Queue, error)
}

type Queue struct {
	Data []int
}

func (q *Queue) Enq(x int) *Queue {
	return &Queue{
		Data: append(q.Data, x),
	}
}

func (q *Queue) Head() (int, error) {
	if q == nil || len(q.Data) == 0 {
		return 0, ErrQueueIsNil
	}
	return q.Data[0], nil
}

func (q *Queue) Deq() (*Queue, error) {
	if q == nil || len(q.Data) == 0 {
		return nil, ErrQueueIsNil
	}
	return &Queue{
		Data: q.Data[1:],
	}, nil
}
