package pointers

import (
	"regexp"
	"strings"
)

/*
Assumptions:
- Arguments: string
- Return: bool
- return true iff the string is a palindrome
- if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters,
    a string reads the same forwards and backwards, then the string is a palindrome
- the above condition is also a necessary condition
- the string consists in only printable ASCII chararcters.

Example:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

*/

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	reg, _ := regexp.Compile("[^a-z0-9]+")
	s = reg.ReplaceAllString(s, "")
	var check func(string) bool
	check = func(s string) bool {
		if len(s) == 0 || len(s) == 1 {
			return true
		}
		if s[0] != s[len(s)-1] {
			return false
		}
		return check(s[1 : len(s)-1])
	}
	return check(s)
}
