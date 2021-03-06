package sol

func longestPalindromeDP(s string) string {
	sLen := len(s)
	resLen := 0
	result := ""
	P := make([][]bool, sLen)
	for row := range P {
		P[row] = make([]bool, sLen)
	}
	for start := sLen - 1; start >= 0; start-- {
		for end := start; end < sLen; end++ {
			if start == end {
				P[start][end] = true
			} else if start == end-1 {
				P[start][end] = s[start] == s[end]
			} else {
				P[start][end] = P[start+1][end-1] && s[start] == s[end]
			}
			if P[start][end] && end-start+1 > resLen {
				resLen = end - start + 1
				result = s[start : end+1]
			}
		}
	}
	return result
}
