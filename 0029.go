func divide(dividend int, divisor int) int {
    sign := 1
	if dividend^divisor < 0 {
		sign = -1
	}

	ldividend, ldivisor := int64(math.Abs(float64(dividend))), int64(math.Abs(float64(divisor)))

	res := 0

	for ldivisor <= ldividend {
		temp := ldivisor
		mul := 1
		for ldividend >= (temp << 1) {
			temp <<= 1
			mul <<= 1
		}
		res += mul
		ldividend -= temp
	}

	res *= sign

	if res >= math.MaxInt32 {
		return math.MaxInt32
	}
	return int(res)
}