package lecture2

const tightStrategyConstant = 4

// GrowableArrayStackTightStrategy ......
type GrowableArrayStackTightStrategy struct {
	_array []interface{}
	t      int
	N      int
}

func createGrowableStackTightStrategy() *GrowableArrayStackTightStrategy {
	return &GrowableArrayStackTightStrategy{t: -1, N: 0}
}

func (stack *GrowableArrayStackTightStrategy) size() int {
	return stack.t + 1
}

func (stack *GrowableArrayStackTightStrategy) top() interface{} {
	if stack.isEmpty() {
		panic("Stack is empty")
	}

	return stack._array[stack.t]
}

func (stack *GrowableArrayStackTightStrategy) isEmpty() bool {
	return stack.t < 0
}

func (stack *GrowableArrayStackTightStrategy) pop() {
	if stack.isEmpty() {
		panic("Stack is empty")
	}

	stack._array[stack.t] = nil
	stack.t--
}

func (stack *GrowableArrayStackTightStrategy) push(data interface{}) {
	if stack.size() == stack.N {
		stack._array = append(stack._array, make([]interface{}, tightStrategyConstant)...)
		stack.N += tightStrategyConstant
	}

	stack.t++
	stack._array[stack.t] = data
}
