package lecture7

// ITree ...
type ITree interface {
	insert(int)
	exists(int) bool
	preOrder() string
	postOrder() string
	inOrder() string
}
