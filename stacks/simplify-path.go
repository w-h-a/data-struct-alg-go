package stacks

import (
	"strings"
)

/*
Assumptions:
- Arguments: string (path)
- Returns: string
- The argument is an absolute path
    - "." refers to the current directory
    - ".." refers to the directory up a level
    - multiple consecutive slashes are treated as "/"
    - anything else is a file/directory name
- the return string is a canonical path
    - the path should start with "/"
    - Any two directorys are separated by a single slash "/"
    - the path does not end with a trailing "/"
    - the path contains and only contains the directories on the path from
        the root dir to the target (no "." or "..")

Examples:

Input: path = "/home/"
Output: "/home"
Explanation: Note that there is no trailing slash after the last directory name.

Input: path = "/../"
Output: "/"
Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.

Input: path = "/home//foo/"
Output: "/home/foo"
Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.
*/

func SimplifyPath(path string) string {
	var helper func([]string, []string) []string
	helper = func(path []string, stack []string) []string {
		if len(path) == 0 {
			return stack
		}
		if path[0] == "" || path[0] == "." {
			return helper(path[1:], stack)
		}
		if path[0] == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			return helper(path[1:], stack)
		}
		return helper(path[1:], append(stack, path[0]))
	}
	result := strings.Join(helper(strings.Split(path, "/"), []string{}), "/")
	return "/" + result
}
