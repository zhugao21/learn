package code

//f(n) = f(n-1) + f(n-2)
func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var a, b = 1, 2
	for i := 3; i <= n; i++ {
		a, b = b%(1e9+7), (a+b)%(1e9+7)
	}
	return b
}
