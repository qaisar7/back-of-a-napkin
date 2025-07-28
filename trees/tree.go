package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	val   int
	left  *node
	right *node
}

var (
	reader *bufio.Reader
)

func init() {
	// Create a new reader that reads from standard input (os.Stdin)
	reader = bufio.NewReader(os.Stdin)

}
func main() {
	values := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, -1, -1, 1}
	n := arrayToTree(values)
	dfs(n)
	newValues := treeToArray(n)
	fmt.Println(truncate(newValues))
	fmt.Println(values)
	for i, v := range values {
		if v != newValues[i] {
			fmt.Printf("value[i]:%d at index %d is not equal to %d", values[i], i, newValues[i])
		}
	}
	fmt.Println(len(newValues))
	fmt.Println(targetSumRecursive(n, 0, 22, []int{}))
	fmt.Println(targetSumIterative(n, 23))
	values = []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4, -1, -1, -1, -1}
	n = arrayToTree(values)

	fmt.Println(leastCommonAncestor(n, 6, 4).val)
}

func leastCommonAncestor(n *node, p, q int) *node {
	fmt.Println("Enter to continue")
	reader.ReadString('\n')

	if n == nil {
		return nil
	}
	fmt.Printf("n.val %d\n", n.val)
	if n.val == p || n.val == q {
		fmt.Printf("left.val || right.val n.val=%d\n", n.val)
		return n
	}
	left := leastCommonAncestor(n.left, p, q)
	right := leastCommonAncestor(n.right, p, q)

	if left != nil && right != nil {
		fmt.Printf("n.val=%d left.val && right.val %d %d\n", n.val, left.val, right.val)
		return n
	}
	if left != nil {
		fmt.Printf("n.val=%d left !=nil %d\n", n.val, left.val)
		return left
	}
	if right != nil {
		fmt.Printf("n.val=%d right.val %d\n", n.val, right.val)
	}
	return right
}

func targetSumIterative(n *node, target int) bool {
	if n == nil {
		return false
	}
	type pair struct {
		n   *node
		sum int
	}
	queue := []pair{{n, n.val}}
	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		if p.n.left != nil {
			queue = append(queue, pair{n: p.n.left, sum: p.n.left.val + p.sum})
		}
		if p.n.right != nil {
			queue = append(queue, pair{n: p.n.right, sum: p.n.right.val + p.sum})
		}
		if p.sum == target {
			return true
		}
	}
	return false
}

func targetSumRecursive(node *node, sum, target int, path []int) (bool, []int) {
	if node == nil {
		return false, path
	}
	sum = node.val + sum
	path = append(path, node.val)
	if sum == target {
		return true, path
	}

	if found, lpath := targetSumRecursive(node.left, sum, target, path); found {
		return true, lpath
	}
	if found, rpath := targetSumRecursive(node.right, sum, target, path); found {
		return true, rpath
	}
	return false, path
}

func treeToArray(n *node) []int {
	// do a bfs
	queue := []*node{n}
	var values []int
	maxLength := 0
	i := 1
	for len(queue) != 0 {
		if len(values) > maxLength {
			break
		}
		n = queue[0]
		queue = queue[1:]
		if n == nil {
			values = append(values, -1)
			queue = append(queue, nil)
			queue = append(queue, nil)
		} else {
			queue = append(queue, n.left)
			queue = append(queue, n.right)
			values = append(values, n.val)
			maxLength = 2*i + 2
		}
		i++
	}
	return values
}

func truncate(values []int) []int {
	last := len(values)
	for i := len(values) - 1; i > 0 && values[i] == -1; i-- {
		last--
	}
	return values[0:last]
}

func dfs(n *node) {
	if n == nil {
		return
	}
	fmt.Println(n.val)
	dfs(n.left)
	dfs(n.right)
}

func arrayToTree(values []int) *node {
	if len(values) == 0 {
		return nil
	}
	mn := map[int]*node{}
	for i := range values {
		var n *node
		var ok bool
		if values[i] == -1 {
			continue
		}
		if n, ok = mn[i]; !ok {
			n = &node{val: values[i]}
			mn[i] = n
		}
		fmt.Printf("n.val %d\n", n.val)
		if i == 5 {
			fmt.Printf("i:%d val:%d left:%d right:%d\n", i, values[i], values[i*2+1], values[i*2+2])
		}
		if i*2+1 < len(values) && values[i*2+1] != -1 {
			lval := &node{val: values[i*2+1]}
			mn[i*2+1] = lval
			n.left = lval
		}
		if i*2+2 < len(values) && values[i*2+2] != -1 {
			rval := &node{val: values[i*2+2]}
			mn[i*2+2] = rval
			n.right = rval
		}
	}
	return mn[0]
}
