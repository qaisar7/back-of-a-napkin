package main

import (
	"fmt"
	"strings"
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

	fmt.Printf("%q\n", urlify(truncate("some   thing    "), "%20"))
	*/

	/* 1.4 - Palindrom Permutation - O(n)
	s := "Tact Csoa"
	fmt.Println(palinPermut(s), palinPermutOptimized(s))
	*/

	/* 1.5 - Find required edits
	strs := [][]string{{"pale", "ple"}, {"palses", "pale"}, {"pale", "bale"}, {"pale", "bake"}}
	for _, s := range strs {
		fmt.Println(checkOneAway(s[0], s[1]))
	}
	*/
	/* 1.6 - Compress string
	 */
	fmt.Println(compressString("aabcccccaaa"))
}

// -------------------1.6----------------------------------
func compressString(s string) string {
	if len(s) <= 1 {
		return s
	}
	var previous byte
	var bytes []byte
	previous = s[0]
	var count int
	for i := 0; i < len(s); i++ {
		if previous == s[i] {
			count++
			continue
		}
		bytes = append(bytes, previous)
		bytes = append(bytes, intToBytes(count)...)
		count = 0
		previous = s[i]
		count++
	}
	bytes = append(bytes, previous)
	bytes = append(bytes, intToBytes(count)...)
	return string(bytes)
}
func intToBytes(i int) []byte {
	s := fmt.Sprintf("%d", i)
	return []byte(s)
}

// -------------------1.5----------------------------------
func checkOneAway(s1, s2 string) bool {
	fmt.Println("comparing ", s1, s2)
	if len(s1)-len(s2) > 1 ||
		len(s2)-len(s1) > 1 {
		return false
	}
	str1 := s1
	str2 := s2
	if len(s1) < len(s2) {
		str1 = s2
		str2 = s1
	}
	i, j := 0, 0
	var mismatch bool
	for i = 0; i < len(str1); i++ {
		if str1[i] != str2[j] {
			if mismatch {
				return false
			}
			mismatch = true
			if len(str1) == len(str2) {
				j++
			}
			continue
		}
		j++
	}
	return true
}

// -------------------1.4----------------------------------
func palinPermut(s string) bool {
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "", -1)
	dict := make(map[byte]int)
	for i := range s {
		if string(s[i]) == " " {
			continue
		}
		dict[s[i]]++
	}
	oneOdd := false
	for _, v := range dict {
		if v%2 == 0 {
			continue
		}
		if len(s)%2 != 0 && !oneOdd {
			oneOdd = true
			continue
		}
		//fmt.Println("because of ", string(k), v, dict)
		return false
	}
	return true
}

func palinPermutOptimized(s string) bool {
	validRange := []byte{}
	a := "a"
	z := "z"
	A := "A"
	Z := "Z"
	validRange = append(validRange, a[0], z[0], A[0], Z[0])
	//fmt.Println(validRange)
	var record uint32
	var diff byte
	length := 0
	for i := range s {
		diff = 0
		// check if the character is between a-z
		if s[i] >= validRange[0] && s[i] <= validRange[1] {
			diff = validRange[0]
		}
		// or is it betwee A-Z ?
		if s[i] >= validRange[2] && s[i] <= validRange[3] {
			diff = validRange[2]
		}
		// if its not between a-z or A-Z then skip
		if diff == 0 {
			continue
		}
		record = record | (uint32(1) << (s[i] - diff))
		//fmt.Printf("length:%d record :%b diff:%b\n", length, record, diff)
		length++
	}
	//fmt.Printf("length:%d record :%b\n", length, record)
	if length%2 == 0 && record != 0 {
		return false
	}
	if length%2 != 0 && !checkSingleOne(record) {
		return false
	}
	return true
}

func checkSingleOne(i uint32) bool {
	j := i - 1
	if j&i == 0 {
		//fmt.Printf("i:%b j:%b\n", i, j)
		return true
	}
	return false
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

// urlify replaces all spaces with the replacement which should typically by '%20'.
func urlify(s string, replacement string) string {
	// One way is to create another slice of bytes with
	// extra space to accomate the replacement.
	var spaces int
	// Count the number of spaces
	for i := range s {
		if string(s[i]) == " " {
			spaces++
		}
	}
	// Create a slice of bytes as big as the string plus extra buffer for the replacement.
	index := len(s) + spaces*(len(replacement)-1)
	bytes := make([]byte, index)
	for i := range s {
		bytes[i] = s[i]
	}
	// Although its not necessary but we are going to start from backwards just to demonstrate that
	// we can manipulate a fixed size slice or array that already has the data.
	// If its a space then copy the replacement.
	// Note the index and 'i' being used separately.
	for i := len(s) - 1; i >= 0; i-- {
		if string(bytes[i]) == " " {
			// copy the replacement in.
			for j := 0; j < len(replacement); j++ {
				bytes[index-len(replacement)+j] = replacement[j]
			}
			index -= len(replacement)
			continue
		}
		// normal copy
		bytes[index-1] = bytes[i]
		index--
	}

	// This is the Go way of easily manipulating a slice, where we won't need anything of the above.
	// We can actually use strings.Replace but that might not be what the interviewer wants to see.
	/*for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			s = s[0:i] + replacement + s[i+1:len(s)]
		}
	}*/
	return string(bytes)
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
