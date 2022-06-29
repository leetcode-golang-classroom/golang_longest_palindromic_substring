package sol

func longestPalindrome(s string) string {
	sLen := len(s)
	resLen := 0
	result := ""
	var checkPalindrome = func(startLeft, startRight int) {
		left := startLeft
		right := startRight
		for left >= 0 && right < sLen && s[left] == s[right] {
			if right-left+1 > resLen {
				resLen = right - left + 1
				result = s[left : right+1]
			}
			left -= 1
			right += 1
		}
	}
	for idx := 0; idx < sLen; idx++ {
		checkPalindrome(idx, idx)
		checkPalindrome(idx, idx+1)
	}
	return result
}
