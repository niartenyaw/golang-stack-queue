package queue

import "github.com/niartenyaw/sequences/stack"

type Queue struct {
	inStack  stack.Stack
	outStack stack.Stack
	length   int
}

func (queue *Queue) Insert(value int) {
	queue.inStack.Push(value)
	queue.updateLength()
}

func (queue *Queue) Remove() int {
	if queue.outStack.Length() == 0 {
		queue.shuffle()
	}

	value := queue.outStack.Pop()
	queue.updateLength()

	return value
}

func (queue *Queue) shuffle() {
	length := queue.inStack.Length()
	for i := 0; i < length; i += 1 {
		queue.outStack.Push(queue.inStack.Pop())
	}
}

func (queue *Queue) updateLength() {
	queue.length = queue.inStack.Length() + queue.outStack.Length()
}
