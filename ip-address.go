package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Assumptions:
- Arguments: string (queryIP)
- Return: string (either "IPv4", "IPv6", "Neither")
- If the argument is a valid IPv4 address, return "IPv4"
- If the argument is a valid IPv6 address, return "IPv6"
- Otherwise, return "Neither"
- IPv4 => "x1.x2.x3.x4",
    where 0 <= x_i <= 255, and
    no leading zeros
- IPv6 => "x1:x2:x3:x4:x5:x6:x7:x8",
    where 1 <= x_1.length <= 4,
    x_i is a hexadecimal string, and
    leading zeros are allowed

Examples:

Input: queryIP = "172.16.254.1"
Output: "IPv4"
Explanation: This is a valid IPv4 address, return "IPv4".

Input: queryIP = "2001:0db8:85a3:0:0:8A2E:0370:7334"
Output: "IPv6"
Explanation: This is a valid IPv6 address, return "IPv6".

Input: queryIP = "256.256.256.256"
Output: "Neither"
Explanation: This is neither a IPv4 address nor a IPv6 address.

*/

func validIPAddress(queryIP string) string {
	if ss := strings.Split(queryIP, "."); len(ss) == 4 {
		if validIPv4(ss) {
			return "IPv4"
		}
	}
	if ss := strings.Split(queryIP, ":"); len(ss) == 8 {
		if validIPv6(ss) {
			return "IPv6"
		}
	}
	return "Neither"
}

func validIPv4(ss []string) bool {
	if len(ss) == 0 {
		return true
	}
	num, err := strconv.Atoi(ss[0])
	if err != nil || num > 255 || num < 0 {
		return false
	}
	if str := fmt.Sprint(num); str != ss[0] {
		return false
	}
	return validIPv4(ss[1:])
}

func validIPv6(ss []string) bool {
	if len(ss) == 0 {
		return true
	}
	str := ss[0]
	if len(str) == 0 || len(str) > 4 {
		return false
	}
	if !validIPv6Bytes(str) {
		return false
	}
	return validIPv6(ss[1:])
}

func validIPv6Bytes(s string) bool {
	if len(s) == 0 {
		return true
	}
	r := s[0]
	if !((r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')) {
		return false
	}
	return validIPv6Bytes(s[1:])
}
