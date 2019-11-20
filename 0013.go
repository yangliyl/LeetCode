func romanToInt(s string) int {
	m := map[string]int{"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
	out := 0
	for i := 1; i < len(s); i++ {
		if m[s[i-1:i]] >= m[s[i:i+1]] {
			out += m[s[i-1:i]]
		} else {
			out -= m[s[i-1:i]]
		}
	}
	out += m[s[len(s)-1:]]
	return out
}
