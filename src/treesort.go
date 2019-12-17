/**
 * 二叉树
 */
package src

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

func Treesort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
		fmt.Println(root)
	}
	apppendValues(values[:0], root)
}

func apppendValues(values []int, t *tree) []int {
	if t != nil {
		values = apppendValues(values, t.left)
		values = append(values, t.value)
		values = apppendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil { // 初始tree
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
