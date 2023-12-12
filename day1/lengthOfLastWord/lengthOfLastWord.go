package main

func lengthOfLastWord(s string) int {
	var ans int

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			ans++
		} else if s[i] == ' ' && ans > 0 {
			return ans
		}
	}

	return ans
}
