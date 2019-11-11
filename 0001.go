func twoSum(nums []int, target int) []int {
    m := map[int]int{}
    for i, v := range nums {
        if k, ok := m[v]; ok {
            return []int{k, i}
        }
        m[target-v] = i
    }
    panic("No two sum solution")
}