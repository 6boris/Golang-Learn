// Package bst implements an unbalanced binary search tree.
package main

// T is the internal representation of a binary search tree.
type T struct {
	root  *node
	count int
}

// node is the internal representation of a binary tree node.
type node struct {
	key  int
	val  interface{}
	l, r *node
}

// TraversalType represents one of the three know traversals.
type TraversalType int

const (
	InOrder TraversalType = iota
	PreOrder
	PostOrder
)

// Insert adds a given key+value to the tree and returns true if it was added.
// Average: O(log(n)) Worst: O(n)
func (t *T) Insert(k int, v interface{}) (added bool) {
	t.root, added = insert(t.root, k, v)
	if added {
		t.count++
	}

	return
}

// insert recusively adds a key+value in the tree.
func insert(n *node, k int, v interface{}) (r *node, added bool) {
	if r = n; n == nil {
		// keep track of how many elements we have in the tree
		// to optimize the channel length during traversal
		r = &node{key: k, val: v}
		added = true
	} else if k < n.key {
		r.l, added = insert(n.l, k, v)
	} else if k > n.key {
		r.r, added = insert(n.r, k, v)
	}

	return
}

// Delete removes a given key from the tree and returns true if it was removed.
// Average: O(log(n)) Worst: O(n)
func (t *T) Delete(k int) (deleted bool) {
	n, deleted := delete(t.root, k)
	if deleted {
		// Handling the case of root deletion.
		if t.root.key == k {
			t.root = n
		}

		t.count--
	}

	return deleted
}

// delete recursively deletes a key from the tree.
func delete(n *node, k int) (r *node, deleted bool) {
	if r = n; n == nil {
		return nil, false
	}

	if k < n.key {
		r.l, deleted = delete(n.l, k)
	} else if k > n.key {
		r.r, deleted = delete(n.r, k)
	} else {
		if n.l != nil && n.r != nil {
			// find the right most element in the left subtree
			s := n.l
			for s.r != nil {
				s = s.r
			}
			r.key = s.key
			r.val = s.val
			r.l, deleted = delete(s, s.key)
		} else if n.l != nil {
			r = n.l
			deleted = true
		} else if n.r != nil {
			r = n.r
			deleted = true
		} else {
			r = nil
			deleted = true
		}
	}

	return
}

// Find returns the value found at the given key.
// Average: O(log(n)) Worst: O(n)
func (t *T) Find(k int) interface{} {
	return find(t.root, k)
}

func find(n *node, k int) interface{} {
	if n == nil {
		return nil
	}

	if n.key == k {
		return n.val
	} else if k < n.key {
		return find(n.l, k)
	} else if k > n.key {
		return find(n.r, k)
	}

	return nil
}

// Clear removes all the nodes from the tree.
// O(n)
func (t *T) Clear() {
	t.root = clear(t.root)
	t.count = 0
}

// clear recursively removes all the nodes.
func clear(n *node) *node {
	if n != nil {
		n.l = clear(n.l)
		n.r = clear(n.r)
	}
	n = nil

	return n
}

// Traverse provides an iterator over the tree.
// O(n)
func (t *T) Traverse(tt TraversalType) <-chan interface{} {
	c := make(chan interface{}, t.count)
	go func() {
		switch tt {

		case InOrder:
			inOrder(t.root, c)
		case PreOrder:
			preOrder(t.root, c)
		case PostOrder:
			postOrder(t.root, c)
		}
		close(c)
	}()

	return c
}

// inOrder returns the left, parent, right nodes.
func inOrder(n *node, c chan interface{}) {
	if n == nil {
		return
	}

	inOrder(n.l, c)
	c <- n.val
	inOrder(n.r, c)
}

// preOrder returns the parent, left, right nodes.
func preOrder(n *node, c chan interface{}) {
	if n == nil {
		return
	}

	c <- n.val
	preOrder(n.l, c)
	preOrder(n.r, c)
}

// postOrder returns the left, right, parent nodes.
func postOrder(n *node, c chan interface{}) {
	if n == nil {
		return
	}

	postOrder(n.l, c)
	postOrder(n.r, c)
	c <- n.val
}
