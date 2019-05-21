package lecture2

// GrowableArrayStackGrowthStrategy ......
type GrowableArrayStackGrowthStrategy struct {
	_array []interface{}
	t      int
	N      int
}

func createGrowableStackGrowthStrategy() *GrowableArrayStackGrowthStrategy {
	return &GrowableArrayStackGrowthStrategy{t: -1, N: 0}
}

func (stack *GrowableArrayStackGrowthStrategy) size() int {
	return stack.t + 1
}

func (stack *GrowableArrayStackGrowthStrategy) top() interface{} {
	if stack.isEmpty() {
		panic("Stack is empty")
	}

	return stack._array[stack.t]
}

func (stack *GrowableArrayStackGrowthStrategy) isEmpty() bool {
	return stack.t < 0
}

func (stack *GrowableArrayStackGrowthStrategy) pop() {
	if stack.isEmpty() {
		panic("Stack is empty")
	}

	stack._array[stack.t] = nil
	stack.t--
}

func (stack *GrowableArrayStackGrowthStrategy) push(data interface{}) {
	if stack.size() == stack.N {
		if stack.size() == 0 {
			stack._array = make([]interface{}, 1)
			stack.N = 1
		} else {
			stack._array = append(stack._array, make([]interface{}, stack.N*2)...)
			stack.N *= 2
		}
	}

	stack.t++
	stack._array[stack.t] = data
}
