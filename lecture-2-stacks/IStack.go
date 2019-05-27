package lecture2

// Stack ...
type Stack interface {
	size() int
	pop()
	push(data interface{})
	top() interface{}
	isEmpty() bool
}
