# golang_longest_palindromic_substring

Given a string `s`, return *the longest palindromic substring* in `s`.

## Examples

**Example 1:**

```
Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.

```

**Example 2:**

```
Input: s = "cbbd"
Output: "bb"

```

**Constraints:**

- `1 <= s.length <= 1000`
- `s` consist of only digits and English letters.

## 解析

題目給定一個字串 s, 要找出字串中最長的回文子字串

如果一個字串 s 是回文 

代表 字串的字元 存在一個 index i, i = floor(len(s)/2)    

使得 s[i+k] == s[i-k] , k = 1 ,…  floor(len(s)/2) if len(s) % 2 == 1

        s[i+k] == s[i-k-1], k=0,…len(s)/2 - 1 if len(s)%2 == 0

來思考一下 回文字串組成要件

“aba” 是回文則 “xabax” 也是回文

也就是要形成回文可以從基本字元開始檢查

解法1:

對所有回文 index

P(i,j) = P(i+1, j-1) && s[i]==s[j] 對所有 j > i

已知 

P(i,i) = true && P(i, i+1) = s[i]==s[j]

這樣只要檢查 從某個 i 開始最大的 j 是多少逐步累加 

對每個 i 只要檢查 j = i+1 到 n 

而只要檢查 n 個開始點 所以時間複雜度是 O($n^2$)

空間複雜度也是 O($n^2$) 因為要儲存每次 P(i,j) 的結果

![](https://i.imgur.com/u6sJCR1.png)
解法2:

可以發現要形成回文, 會有一個中間點，其他兩邊是相同的

從每個字元當作中間字元開始往兩邊檢查

要檢查兩種情況

1 回文長度是奇數 代表中間點是一個字元

2 回文長度是偶數 代表終點點是兩個相同字元

每次往外檢查到邊緣 O(n)

要檢查 n 個中間值

所以時間複雜度是 O($n^2$)

因為不用特別去紀錄運算過的結果 所以空間複雜度 O(1)

![](https://i.imgur.com/Inl8a70.png)

## 程式碼
```go
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
```
## 困難點

1. 要找出 palindrome 的遞迴結構

## Solve Point

- [x]  透過 palindrome 特性由長度小到大檢查該 substring 是否是 palindrome
- [x]  需要檢查 odd / even length 兩種可能的 pattern