func reverse(x int) int {
	sum := 0
	for x != 0 {
		sum = sum*10 + x%10
		x /= 10
		if sum > math.MaxInt32 || sum < math.MinInt32 {
			return 0
		}
	}
	return sum
}
