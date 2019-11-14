func longestPalindrome(s string) string {
	if len(s) < 1 {
		return s
	}

	start, max := 0, 0

	for i := 0; i < len(s); i++ {
		extendPalindrome := func(left, right int) {
			for left >= 0 && right < len(s) && s[left] == s[right] {
				left --
				right ++
			}
		
			if max < right - left - 1 {
				start = left + 1
				max = right - left - 1
			}
		}
		extendPalindrome(i, i)
		extendPalindrome(i, i+1)
	}

	return s[start:start + max]
}
