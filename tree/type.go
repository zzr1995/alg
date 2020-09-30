package tree

// node
type node struct {
	data        int
	left, right *node
	parent      *node
}

type bst struct {
	root   *node
	length int
}
