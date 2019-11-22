func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	out := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			return out
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l := i + 1
		r := len(nums) - 1
		for l < r {
			tmp := nums[i] + nums[l] + nums[r]
			if tmp > 0 {
				r--
			} else if tmp < 0 {
				l++
			} else {
				out = append(out, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}
	return out
}
