package lecture3

// Queue ...
type Queue interface {
	size() int
	front() int
	enqueue(int) bool
	dequeue() bool
}
