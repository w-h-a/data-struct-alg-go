package stacks

import (
	"errors"

	"github.com/w-h-a/more-dsa/lists"
)

var (
	ErrStackIsNil = errors.New("stack is nil")
)

type Stack struct {
	Data *lists.ListNode
}

func Push(s *Stack, x int) *Stack {
	return &Stack{
		Data: &lists.ListNode{
			Val:  x,
			Next: s.Data.Next,
		},
	}
}

func Top(s *Stack) (int, error) {
	if s == nil || s.Data == nil {
		return 0, ErrStackIsNil
	}
	return s.Data.Val, nil
}

func Pop(s *Stack) (*Stack, error) {
	if s == nil || s.Data == nil {
		return nil, ErrStackIsNil
	}
	return &Stack{
		Data: s.Data.Next,
	}, nil
}
