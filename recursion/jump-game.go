package recursion

/*
Assumptions:
- Arguments: int slice (xs)
- Returns: bool
- Each element represents your _max_ jump length
- Return true if and only if you can reach the last index

Examples:

Input: nums = [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

The second one is due to Trevor Kelly
The third one is due to Josh Leath
*/

func CanJump(xs []int) bool {
	cache := map[int]bool{}
	var memoized func(int) bool
	memoized = func(idx int) bool {
		if idx == len(xs)-1 {
			return true
		}
		key := idx
		_, ok := cache[key]
		if !ok {
			val := xs[idx]
			if val >= len(xs)-idx {
				cache[key] = true
			} else {
				for val > 0 {
					cache[key] = memoized(idx + val)
					if cache[key] {
						return cache[key]
					}
					val--
				}
			}
		}
		return cache[key]
	}
	return memoized(0)
}

func CanJumpII(nums []int) bool {
	pointer, maxReach := 0, 0
	for maxReach < len(nums)-1 {
		reach := pointer + nums[pointer]
		if reach > maxReach {
			maxReach = reach
		}
		if maxReach <= pointer {
			return false
		}
		pointer++
	}
	return true
}

func CanJumpIII(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	lastPossible := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if lastPossible <= i+nums[i] {
			lastPossible = i
		}
	}
	return lastPossible == 0
}
