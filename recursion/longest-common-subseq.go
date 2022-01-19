package recursion

/*
Assumptions:
- Arguments: string (text1), string (text2)
- Returns: int
- Return the length of their longest common subsequence.
- A subsequence of a string is a new string generated from the original with 0 or more charcters deleted without changing order
- A common subsequence of two string is a subsequence that is common to both strings

Examples:

Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.

Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.

Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.

the first solution below is Josh Leath's solution
the second solution is mine
*/

func LongestCommonSubsequence(text1 string, text2 string) int {
	dp := make(map[[2]int]int)

	for i, c1 := range text1 {
		for j, c2 := range text2 {
			key := [2]int{i, j}
			left := dp[[2]int{i, j - 1}]
			up := dp[[2]int{i - 1, j}]
			diagonal := dp[[2]int{i - 1, j - 1}]
			if c1 != c2 {
				dp[key] = max(left, up)
			} else {
				dp[key] = diagonal + 1
			}
		}
	}
	return dp[[2]int{len(text1) - 1, len(text2) - 1}]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func LongestCommonSubsequenceRec(text1, text2 string) int {
	cache := map[[2]string]int{}
	var helper func(string, string) int
	helper = func(x, y string) int {
		if len(x) == 0 || len(y) == 0 {
			return 0
		}
		key := [2]string{x, y}
		_, ok := cache[key]
		if !ok {
			if x[0] == y[0] {
				cache[key] = 1 + helper(x[1:], y[1:])
			} else {
				cache[key] = max(helper(x[1:], y), helper(x, y[1:]))
			}
		}
		return cache[key]
	}
	return helper(text1, text2)
}
