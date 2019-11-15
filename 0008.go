func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	begin := false
	sign := false
	out := 0
	for _, v := range str {
		if !begin {
			if v == '+' {
				begin = true
				continue
			} else if v == '-' {
				begin = true
				sign = true
				continue
			} else if v == ' ' {
				continue
			}
		}
		if v >= '0' && v <= '9' {
			begin = true
			if !sign && (out > math.MaxInt32/10 || (out == math.MaxInt32/10 && v-'0' > 7)) {
				return math.MaxInt32
			} else if sign && (out > math.MaxInt32/10 || (out == math.MaxInt32/10 && v-'0' > 7)) {
				return math.MinInt32
			} else {
				out = out*10 + int(v-'0')
				continue
			}
		}
		break
	}

	if sign {
		out *= -1
	}

	return out
}
