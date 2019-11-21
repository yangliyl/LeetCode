func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) < 1 {
		return ""
	}

	pre := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], pre) != 0 {
			pre = pre[0 : len(pre)-1]
		}
	}
	return pre
}
