package stacks

/*
Assumptions:
- Arguments: string (s)
- Returns: bool
- The string contains only '(', ')', '{', '}', '[' and ']'
- Return true if and only if
    - open brackets are closed by the same type of bracket, and
    - open brackets are closed in the correct order

Examples:

Input: s = "()"
Output: true

Input: s = "()[]{}"
Output: true

Input: s = "(]"
Output: false
*/

func IsValid(s string) bool {
	hashMap := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	stack := []string{}
	var solution func(string) bool
	solution = func(s string) bool {
		if len(s) == 0 {
			return len(stack) == 0
		}
		_, ok := hashMap[string(s[0])]
		if ok {
			stack = append(stack, string(s[0]))
		} else {
			if len(stack) > 0 {
				key := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if hashMap[key] != string(s[0]) {
					return false
				}
			} else {
				return false
			}
		}
		return solution(s[1:])
	}
	return solution(s)
}
