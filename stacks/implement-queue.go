package stacks

type MyQueue struct {
	Heads *Stack
	Tails *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		Heads: &Stack{},
		Tails: &Stack{},
	}
}

// Enq
func (this *MyQueue) Push(x int) {
	this.Tails = Push(this.Tails, x)
	normalize(this)
}

// Deq
func (this *MyQueue) Pop() int {
	x, _ := Top(this.Heads)
	heads, _ := Pop(this.Heads)
	this.Heads = heads
	normalize(this)
	return x
}

// Head
func (this *MyQueue) Peek() int {
	x, _ := Top(this.Heads)
	return x
}

// IsEmpty
func (this *MyQueue) Empty() bool {
	return IsEmpty(this.Heads)
}

func normalize(q *MyQueue) {
	if q.Heads == nil || q.Heads.Data == nil {
		q.Heads = remove(&Stack{}, q.Tails)
		q.Tails = &Stack{}
	}
}

func remove(heads, tails *Stack) *Stack {
	if tails.Data == nil {
		return heads
	}
	x, _ := Top(tails)
	tails, _ = Pop(tails)
	return remove(Push(heads, x), tails)
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
