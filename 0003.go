func lengthOfLongestSubstring(s string) int {
    m := map[byte]int{}
    count := 0
    for i, j := 0, 0; j < len(s); j++ {
        if t, ok := m[s[j]]; ok {
            if t > i {
                i = t
            }
        }
        sum := j - i + 1
        if count < sum {
            count = sum
        }
        m[s[j]] = j + 1
    }
    return count
}