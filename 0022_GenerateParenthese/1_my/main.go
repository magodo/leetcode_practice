package main

import "fmt"

type Node struct {
	val            Value
	left           *Node
	right          *Node
	remainingOpen  int
	remainingClose int
}

type Value string

const (
	openParenthes  Value = "("
	closeParenthes Value = ")"
)

func newTree(n int) *Node {
	if n == 0 {
		return nil
	}
	root := &Node{
		val:            openParenthes,
		remainingOpen:  n - 1,
		remainingClose: n,
	}
	root.generateChildren()
	return root
}

func (n *Node) generateChildren() {
	if n.remainingOpen == 0 && n.remainingClose == 0 {
		return
	}

	if n.remainingClose > n.remainingOpen {
		n.left = &Node{
			val:            closeParenthes,
			remainingOpen:  n.remainingOpen,
			remainingClose: n.remainingClose - 1,
		}
		n.left.generateChildren()
	}

	if n.remainingOpen > 0 {
		n.right = &Node{
			val:            openParenthes,
			remainingOpen:  n.remainingOpen - 1,
			remainingClose: n.remainingClose,
		}
		n.right.generateChildren()
	}
}

func (n *Node) traverse(s string, out *[]string) {
	s = fmt.Sprintf("%s%s", s, n.val)
	if n.left == nil && n.right == nil {
		*out = append(*out, s)
	}
	if n.left != nil {
		n.left.traverse(s, out)
	}
	if n.right != nil {
		n.right.traverse(s, out)
	}
}

func generateParenthesis(n int) []string {
	root := newTree(n)
	if root == nil {
		return nil
	}
	out := []string{}
	root.traverse("", &out)
	return out
}
