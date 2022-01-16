package recursion

/*
Assum:
- Arguments: float (x), int (n)
- Returns: float
- Calculate x^n
- Should work for n = 0 and n < 0

Examples:

Input: x = 2.00000, n = 10
Output: 1024.00000

Input: x = 2.10000, n = 3
Output: 9.26100

Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2^-2 = 1/2^2 = 1/4 = 0.25

Main idea:
- Brute force: Multiply x until n is 0
    - O(n)
    - runtime errors for large n
- Another idea: half n each run and do the necessary calculations
    based on whether n is even or odd
    - for example, let x = 3.0, n = 20
    - 3^20
        => (3^2)^10
            (i.e., if n is even, feed x*x and n/2 back in)
        => (9^2)^5
            (same thing as before)
        => (81 * 81^4)
            (math)
        => ( 81 * (81^2)^2 )
            (i.e., moving from n = 5 => n = 2 requires that we feed x*x and n/2 back in but also make sure to
            multiply that eventual return value by another x)
    - O(log n)
    - To make it work for n < 0, feed function x and -n. Once we get the result, return 1 / result
    - n = 0 is a special case; simply return 1
*/

func MyPow(x float64, k int) float64 {
	switch {
	case k < 0:
		return 1 / MyPow(x, -k)
	case k == 0:
		return 1
	default:
		if k == 1 {
			return x
		}
		if k%2 == 0 {
			return MyPow(x*x, k/2)
		}
		return x * MyPow(x*x, k/2)
	}
}
