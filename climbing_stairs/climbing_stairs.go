/*
Climbing Stairs
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:

Input: n = 2
Output: 2
Explanation:	 There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

Constraints:

	1 <= n <= 45
*/
package main

import "fmt"

func climbStairs(n int) int {
	memo := make(map[int]int)
	var cal func(int) int
	cal = func(n int) int {
		if value, exits := memo[n]; exits {
			return value
		}
		if n == 1 {
			return 1 // step 1
		}
		if n == 2 {
			return 2 // step 2
		}
		memo[n] = cal(n-1) + cal(n-2)
		return memo[n]
	}
	return cal(n)
}

func main() {
	number := 45
	result := climbStairs(number)

	fmt.Println(result)
}
