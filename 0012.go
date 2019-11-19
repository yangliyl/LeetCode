func intToRoman(num int) string {
	nums := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	buf := bytes.Buffer{}
	for i := 0; i < len(nums); i++ {
		if num >= nums[i] {
			buf.WriteString(strings.Repeat(romans[i], num/nums[i]))
			num %= nums[i]
		}
	}
	return buf.String()
}
