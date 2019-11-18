package main

import "fmt"

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(a))
}

func maxArea(height []int) int {
	i, j, max := 0, len(height)-1, 0
	for i < j {
		tmp := 0
		if height[i] < height[j] {
			tmp = height[i] * (j - i)
			i++
		} else {
			tmp = height[j] * (j - i)
			j--
		}

		if max < tmp {
			max = tmp
		}
	}
	return max
}
