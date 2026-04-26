// Level 1: Valid Anagram
//
// Task: Write the logic for `isAnagram`. It takes two strings, `s` and `t`.
// It should return `true` if `t` is an anagram of `s`, and `false` otherwise.
// An Anagram is a word formed by rearranging the letters of a different word,
// using all the original letters exactly once.
//
// Constraints:
// - Assume the strings contain only lowercase English letters.
// - Try to make it efficient (O(N) time complexity) using Go maps or slices.

package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var counts [26]int

	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}
	for _, count := range counts {

		if count != 0 {
			return false

		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Printf("Test 1 (true): %v\n", isAnagram("anagram", "nagaram"))
	fmt.Printf("Test 2 (false): %v\n", isAnagram("rat", "car"))
	fmt.Printf("Test 3 (true): %v\n", isAnagram("listen", "silent"))
	fmt.Printf("Test 4 (false): %v\n", isAnagram("a", "ab"))
}
