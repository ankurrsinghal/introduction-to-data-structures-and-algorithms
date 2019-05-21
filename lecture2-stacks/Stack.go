package lecture2

const capacity = 1024

// ArrayStack ...
type ArrayStack struct {
	_array []interface{}
	t      int
	N      int
}

// Non Growable Stack
func createStack(cap int) *ArrayStack {
	return &ArrayStack{_array: make([]interface{}, cap), t: -1, N: cap}
}

func (stack *ArrayStack) size() int {
	return stack.t + 1
}

func (stack *ArrayStack) top() interface{} {
	if stack.isEmpty() {
		panic("Stack is empty")
	}

	return stack._array[stack.t]
}

func (stack *ArrayStack) isEmpty() bool {
	return stack.t < 0
}

func (stack *ArrayStack) pop() {
	if stack.isEmpty() {
		panic("Stack is empty")
	}

	stack._array[stack.t] = nil
	stack.t--
}

func (stack *ArrayStack) push(data interface{}) {
	if stack.size() == stack.N {
		panic("StackOverflow")
	}

	stack.t++
	stack._array[stack.t] = data
}
