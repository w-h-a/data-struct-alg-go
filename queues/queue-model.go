package queues

import (
	"errors"

	"github.com/w-h-a/more-dsa/lists"
)

var (
	ErrQueueIsNil = errors.New("queue is nil")
)

type Queue struct {
	Heads *lists.ListNode
	Tails *lists.ListNode
}

func Enq(q *Queue, x int) *Queue {
	return normalize(&Queue{
		Heads: q.Heads,
		Tails: &lists.ListNode{
			Val:  x,
			Next: q.Tails,
		},
	})
}

func Head(q *Queue) (int, error) {
	if q == nil || q.Heads == nil {
		return 0, ErrQueueIsNil
	}
	return q.Heads.Val, nil
}

func Deq(q *Queue) (*Queue, error) {
	if q == nil || q.Heads == nil {
		return nil, ErrQueueIsNil
	}
	return normalize(&Queue{
		Heads: q.Heads.Next,
		Tails: q.Tails,
	}), nil
}

func IsEmpty(q *Queue) bool {
	return q.Heads == nil
}

func normalize(q *Queue) *Queue {
	if q.Heads == nil {
		return &Queue{
			Heads: lists.Reverse(q.Tails),
		}
	}
	return q
}
