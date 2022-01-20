package queues

type MyStack struct {
	Popper *Queue
	Pusher *Queue
}

func Constructor() MyStack {
	return MyStack{
		Popper: &Queue{},
		Pusher: &Queue{},
	}
}

func (this *MyStack) Push(x int) {
	this.Pusher = Enq(this.Pusher, x)
	fix(this)
}

func (this *MyStack) Pop() int {
	x, _ := Head(this.Popper)
	popper, _ := Deq(this.Popper)
	this.Popper = popper
	return x
}

func (this *MyStack) Top() int {
	x, _ := Head(this.Popper)
	return x
}

func (this *MyStack) Empty() bool {
	return IsEmpty(this.Popper)
}

func fix(s *MyStack) {
	pusher := remove(s.Pusher, s.Popper)
	s.Pusher = &Queue{}
	s.Popper = pusher
}

func remove(pusher, popper *Queue) *Queue {
	if popper.Heads == nil {
		return pusher
	}
	x, _ := Head(popper)
	popper, _ = Deq(popper)
	return remove(Enq(pusher, x), popper)
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
