package main

import (
	"fmt"
)

type node struct {
	val   int
	left  *node
	right *node
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
	fmt.Println(leastCommonAncestor(n, 6, 8))
}

func leastCommonAncestor(n *node, p, q int) ([]int, []int, int) {
	// find the common ancestor of p and q. pA denotes the ancestor so far found for p and qA for q.
	if n == nil {
		return []int{}, []int{}, -1
	}
	pAL, qAL, aL := leastCommonAncestor(n.left, p, q)
	pAR, qAR, aR := leastCommonAncestor(n.right, p, q)
	pA, qA := []int{}, []int{}
	a := -1
	if len(pAL) != 0 {
		pA = pAL
	}
	if len(pAR) != 0 {
		pA = pAR
	}
	if len(qAL) != 0 {
		qA = qAL
	}
	if len(qAR) != 0 {
		qA = qAR
	}
	if aL != -1 {
		a = aL
	}
	if aR != -1 {
		a = aR
	}
	if n.val == p {
		pA = append(pA, p)
		return pA, qA, a
	}
	if n.val == q {
		qA = append(qA, q)
		return pA, qA, a
	}
	if len(pA) != 0 {
		pA = append(pA, n.val)
	}
	if len(qA) != 0 {
		qA = append(qA, n.val)
	}
	if a == -1 && len(pA) != 0 && len(qA) != 0 {
		pl := len(pA)
		ql := len(qA)
		var shortOne []int
		var longOne []int
		shortOne = qA
		longOne = pA
		if pl < ql {
			shortOne = pA
			longOne = qA
		}
		keys := map[int]bool{}
		for i := 0; i < len(longOne); i++ {
			keys[longOne[i]] = true
		}
		for i := 0; i < len(shortOne); i++ {
			if _, ok := keys[shortOne[i]]; ok {
				a = shortOne[i]
				return pA, qA, a
			}
		}
	}
	fmt.Println(pA, qA, a)
	return pA, qA, a
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
