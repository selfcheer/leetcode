package dynamic_programming

import "fmt"

func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	dp :=  make([]int, n + 1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i ++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func TestClimbStairs() {
	fmt.Println(climbStairs(1))
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
	fmt.Println(climbStairs(5))
}
