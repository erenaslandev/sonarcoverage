package sonarcoverage

var database = make(map[int]int)

func FibonacciWithoutCache(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciWithoutCache(n-1) + FibonacciWithoutCache(n-2)
}

func FibonacciWithCache(n int) int {
	if n <= 1 {
		return n
	}
	if result, ok := database[n]; ok {
		return result
	}
	result := FibonacciWithCache(n-1) + FibonacciWithCache(n-2)
	database[n] = result
	return result
}
