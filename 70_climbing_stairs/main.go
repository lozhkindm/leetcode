package main

func main() {

}

func climbStairs(n int) int {
	prev1 := 1
	prev2 := 0
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			prev2 += prev1
		} else {
			prev1 += prev2
		}
	}
	return prev1 + prev2
}
