package main

import "fmt"

func maxArea(height []int) int {
	largestArea := 0
	left := 0
	right := len(height) - 1
	for left < right {
		area := right - left
		if height[left] > height[right] {
			area *= height[right]
			right--
		} else {
			area *= height[left]
			left++
		}
		if area > largestArea {
			largestArea = area
		}
	}

	return largestArea
}

func main() {
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}
