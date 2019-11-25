package main

import "math"

import "fmt"

import "sort"

func main() {
	arr := []int{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Println(threeSumClosest(arr, 82))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	rs := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-rs)) {
				rs = sum
			}
			if sum < target {
				l++
			} else if sum > target {
				r--
			} else {
				return rs
			}
		}
	}
	return rs
}
