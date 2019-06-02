package lecture7

import (
	"fmt"
	"testing"
)

func TestPreOrder(t *testing.T) {
	tree := TreeConstructor()
	tree.insert(10, 0)
	tree.insert(10, 0)
	tree.insert(1, 0)
	tree.insert(4, 0)
	tree.insert(6, 0)
	tree.insert(9, 0)
	tree.insert(16, 0)

	fmt.Println(tree.exists(10))

	tree.inOrderWalk()
	fmt.Println("Hello")
	tree.preOrderWalk()
	fmt.Println("Hello")
	tree.postOrderWalk()
}
