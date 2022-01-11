package main

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

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	reg, _ := regexp.Compile("[^a-z0-9]+")
	s = reg.ReplaceAllString(s, "")
	i := 0
	j := len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
