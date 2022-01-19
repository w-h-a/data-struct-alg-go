package stacks

import (
	"errors"
)

var (
	ErrStackIsNil = errors.New("stack is nil")
)

/*
type Stacker interface {
    Push(x int) *Stack
    Top() (int, error)
    Pop() (*Stack, error)
}

type Stack struct {
    Data []int
}

func (s *Stack) Push(x int) *Stack {
    return &Stack{
        Data: append([]int{x}, s.Data...),
    }
}

func (s *Stack) Top() (int, error) {
    if s == nil || len(s.Data) == 0 {
        return 0, ErrStackIsNil
    }
    return s.Data[0], nil
}

func (s *Stack) Pop() (*Stack, error) {
    if s == nil || len(s.Data) == 0 {
        return nil, ErrStackIsNil
    }
    return &Stack{
        Data: s.Data[1:],
    }, nil
}
*/

type Stack struct {
	Data []int
}

func Push(s *Stack, x int) *Stack {
	return &Stack{
		Data: append([]int{x}, s.Data...),
	}
}

func Top(s *Stack) (int, error) {
	if s == nil || len(s.Data) == 0 {
		return 0, ErrStackIsNil
	}
	return s.Data[0], nil
}

func Pop(s *Stack) (*Stack, error) {
	if s == nil || len(s.Data) == 0 {
		return nil, ErrStackIsNil
	}
	return &Stack{
		Data: s.Data[1:],
	}, nil
}
