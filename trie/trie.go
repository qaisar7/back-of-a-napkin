package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type node struct {
	character byte
	count     int32
	children  map[byte]*node
}

func insert(trie *node, str string) {
	if trie == nil {
		panic(fmt.Sprintf("empty node provided for %s", str))
	}
	if len(str) == 0 {
		return
	}
	if trie.children == nil {
		trie.children = make(map[byte]*node)

	}
	r := str[0]
	if trie.children[r] == nil {
		trie.children[r] = &node{
			character: r,
			count:     0,
		}
	}
	trie.children[r].count++
	insert(trie.children[r], str[1:len(str)])
}

func find(nodes map[byte]*node, str string) int32 {
	if n, ok := nodes[str[0]]; ok {
		if n.count != 0 {
			if len(str) == 1 {
				return n.count
			}
			return find(n.children, str[1:len(str)])
		}
		return 0
	}
	return 0
}

/*
 * Complete the contacts function below.
 */
func contacts(queries [][]string) []int32 {
	root := &node{}
	results := []int32{}
	for _, q := range queries {
		if q[0] == "add" {
			insert(root, q[1])
		}
		if q[0] == "find" {
			results = append(results, find(root.children, q[1]))
		}
	}
	return results
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	queriesRows, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var queries [][]string
	for queriesRowItr := 0; queriesRowItr < int(queriesRows); queriesRowItr++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []string
		for _, queriesRowItem := range queriesRowTemp {
			queriesRow = append(queriesRow, queriesRowItem)
		}

		if len(queriesRow) != int(2) {
			fmt.Println(queriesRow)
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := contacts(queries)

	for resultItr, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if resultItr != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()

	print()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
