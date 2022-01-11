package main

/*
Assumptions:
- Arguments: needle, haystack string
- Return: integer
- if needle is not part of haystack, return -1
- if needle is part of haystack, return index of first occurrence of needle in haystack
- needle is part of haystack just in case there is a substring in haystack that is identical to needle
- the index of the first occurrence of needle is the index where the first occurrence of needle begins

Examples:

Input: haystack = "hello", needle = "ll"
Output: 2

Input: haystack = "aaaaa", needle = "bba"
Output: -1

Input: haystack = "", needle = ""
Output: 0

*/

func strStr1(haystack string, needle string) int {
	for i := 0; ; i++ {
		for j := 0; ; j++ {
			if j == len(needle) {
				return i
			}
			if i+j == len(haystack) {
				return -1
			}
			if needle[j] != haystack[i+j] {
				break
			}
		}
	}
}

func strStr2(haystack, needle string) int {
	lenNeedle := len(needle)
	lenHaystk := len(haystack)
	for i := 0; i <= lenHaystk-lenNeedle; i++ {
		if haystack[i:i+lenNeedle] == needle {
			return i
		}
	}
	return -1
}
