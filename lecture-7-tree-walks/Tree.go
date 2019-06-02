package lecture7

import "fmt"

// Node ...
type Node struct {
	parent *Node
	left   *Node
	right  *Node
	key    int
	data   interface{}
}

// Tree ...
type Tree struct {
	root *Node
	size int
}

// TreeConstructor ...
func TreeConstructor() Tree {
	return Tree{root: nil, size: 0}
}

func (tree *Tree) insert(key int, data interface{}) {
	node := Node{key: key, data: data}
	if tree.root == nil {
		tree.root = &node
	} else {
		var p *Node
		for i := tree.root; i != nil; {
			p = i
			if i.key < key {
				i = i.right
			} else {
				i = i.left
			}
		}
		if p.key < key {
			p.right = &node
		} else {
			p.left = &node
		}
		node.parent = p
	}
}

func (tree *Tree) exists(key int) bool {
	return tree.search(key) != nil
}

func (tree *Tree) search(key int) *Node {
	var i *Node
	for i = tree.root; i != nil; {
		if i.key == key {
			break
		} else if i.key < key {
			i = i.right
		} else {
			i = i.left
		}
	}
	return i
}

func (tree *Tree) delete(key int) {
	node := tree.search(key)
	if node.left == nil && node.right == nil {
		if node.parent.left == node {
			node.parent.left = nil
		} else {
			node.parent.right = nil
		}
		node.parent = nil
	} else if node.left == nil || node.right == nil {

	} else {

	}
}

func preOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.key)
	preOrder(node.left)
	preOrder(node.right)
}

func postOrder(node *Node) {
	if node == nil {
		return
	}

	postOrder(node.left)
	postOrder(node.right)
	fmt.Println(node.key)
}

func inOrder(node *Node) {
	if node == nil {
		return
	}

	inOrder(node.left)
	fmt.Println(node.key)
	inOrder(node.right)
}

func (tree *Tree) inOrderWalk() {
	inOrder(tree.root)
}

func (tree *Tree) preOrderWalk() {
	preOrder(tree.root)
}

func (tree *Tree) postOrderWalk() {
	postOrder(tree.root)
}
