package slidingwindow

import (
	"slices"
)

func SubstringAnagram(s string, t string) int {
	/*
		Anagrams -> can check against with each occurring frequencies

		Time Complexity: O(n)
		Space Complexity: O(1)
	*/

	count := 0
	first, second := 0, 0
	expectedFreqs, freqs := make([]int, 26), make([]int, 26)
	for _, c := range t {
		expectedFreqs[c-'a'] += 1
	}
	for second < len(s) {
		freqs[s[second]-'a'] += 1
		if second-first+1 == len(t) {
			if slices.Equal(freqs, expectedFreqs) {
				count += 1
			}
			freqs[s[first]-'a'] -= 1
			first += 1
		}
		second += 1
	}
	return count
}
