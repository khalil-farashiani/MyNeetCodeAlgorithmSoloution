package main

func isAnagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}
	letterFreq := make(map[rune]int, len(s))

	for index, letter := range s {
		letterFreq[letter]++
		letterFreq[rune(t[index])]--
	}
	for _, val := range letterFreq {
		if val != 0 {
			return false
		}
	}
	return true
}

// https://leetcode.com/problems/valid-anagram/description/
func main() {
	// testCase
	println(isAnagram("rat", "car"))
}
