package pointers

/*
Assumptions:
- Arguments: two strings (s and t)
- Returns: bool
- return true iff s is a subsequence of t
- s is a subsequence of t just in case
    - s is formed from t by deleting 0 or more characters from t
    - without disturbing relative positions of remaining characters

Questions:
What should I do if len(s) == 0?
    - Presumed response:
        - return true ("" is, indeed, formed by deleting from t 0 or more characters without distruption)

Examples:

Input: s = "abc", t = "ahbgdc"
Output: true

Input: s = "axc", t = "ahbgdc"
Output: false
*/

func IsSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}
	if s[0] == t[0] {
		return IsSubsequence(s[1:], t[1:])
	}
	return IsSubsequence(s, t[1:])
}
