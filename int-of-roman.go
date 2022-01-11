package main

/*
Assumptions:
- Arguments: string
- Return: int
- The length of the string will be between 1 and 15, inclusive
- The integer will be between 1 and 3999, inclusive
- Roman numerals are represented by 'I', 'V', 'X', 'L', 'C', 'D', and 'M':
    'I' => 1
    'V' => 5
    'X' => 10
    'L' => 50
    'C' => 100
    'D' => 500
    'M' => 1000
- Roman numerals are usually written largest to smallest from left to right, in which case they are added together
- When a smaller one comes before a larger one,
    total = total + current integer - 2 * previous integer
    e.g.: IV
        'I' => total = 1
        'V' => total = 1 + 5 - 2 * 1 = 4
    e.g.: XL
        'X' => total = 10
        'L' => total = 10 + 50 - 2 * 10 = 40

Examples:

Input: s = "III"
Output: 3
Explanation: III = 3.

Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.

Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

*/

func romanToInt(s string) int {
	var intOfRom func(int, int, string) int
	intOfRom = func(prev, total int, s string) int {
		if len(s) == 0 {
			return total
		}
		curr := hash[s[0]]
		if curr > prev {
			total = total + curr - 2*prev
		} else {
			total = total + curr
		}
		return intOfRom(curr, total, s[1:])
	}
	return intOfRom(0, 0, s)
}

var hash = map[byte]int{
	'I': 1, 'V': 5, 'X': 10,
	'L': 50, 'C': 100, 'D': 500,
	'M': 1000,
}
