package recursion

/*
Assumptions:
- Arguments: int slice (coin denoms), int (amount)
- Returns: int
- If the change (list of coins) can be made, return number representing fewest number of coins needed
- Otherwise, return -1
- You may assume that you have an infinite number of each kind of coin

Examples:

Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Input: coins = [2], amount = 3
Output: -1

Input: coins = [1], amount = 0
Output: 0

the answer below times out on leetcode
*/

func CoinChange(coins []int, amount int) int {
	result := allWays([][]int{}, []int{}, coins, amount)
	if len(result) == 0 {
		return -1
	}
	return len(result[0])
}

func allWays(result [][]int, soFar, coinVals []int, amount int) [][]int {
	if len(coinVals) == 0 || amount < 0 {
		return result
	}
	if amount == 0 {
		if len(result) == 0 {
			return append(result, soFar)
		}
		if len(soFar) < len(result[0]) {
			result = [][]int{}
			return append(result, soFar)
		}
		return result
	}
	c := coinVals[0]
	return allWays(allWays(result, soFar, coinVals[1:], amount), append([]int{c}, soFar...), coinVals, amount-c)
}
