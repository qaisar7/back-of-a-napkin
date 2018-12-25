package main

import (
	"fmt"
)

func main() {
	/* 1.1  - find if all characters are unique in given string : O(n)
	a := "abcdefghijklmnopqrstuvwxyz  "
	isUnique(a)
	*/

	/* 1.2 - find if one string is a permuation of another : O(s + t)
	s := "somehowb"
	t := "ewmoosha"
	fmt.Println(checkPermuation(s, t))
	*/

	/* 1.3 - URLify
	 */
	fmt.Printf("%q\n", urlify(truncate("some   thing    ")))

	fmt.Println("hello")

}

// -------------------1.3----------------------------------
func truncate(s string) string {
	for i := len(s) - 1; i > 0; i-- {
		if string(s[i]) == " " {
			s = s[0:i]
			continue
		}
		break
	}
	return s
}
func urlify(s string) string {
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			s = s[0:i] + "%20" + s[i+1:len(s)]
		}
	}
	return s
}

// -------------------1.1----------------------------------
func isUnique(str string) bool {
	var check uint32
	for _, v := range str {
		if check&(1<<(uint32(v)-uint32('a'))) > 0 {
			fmt.Printf("character %s is not unique\n", string(v))
			return false
		}
		check = check | (1 << (uint32(v) - uint32('a')))
	}
	return true
}

// --------------------1.2---------------------------------
func checkPermuation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1Count := characterCount(s1)
	s2Count := characterCount(s2)
	fmt.Println(s1Count, s2Count)
	for k, v := range s1Count {
		if v != s2Count[k] {
			return false
		}
	}
	return true
}

func characterCount(s string) map[rune]int {
	cc := make(map[rune]int)
	for _, v := range s {
		cc[v]++
	}
	return cc
}

// -----------------------------------------------------
