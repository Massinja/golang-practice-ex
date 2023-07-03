package main

import "fmt"

func maxArea(height []int) int {
	largestArea := 0
	lastIndex := len(height) - 1
	for i, val := range height[:lastIndex] {
		for k2, valK2 := range height[i+1:] {
			wall := val
			if wall > valK2 {
				wall = valK2
			}
			area := (k2 + 1) * wall
			if largestArea < area {
				largestArea = area
			}
		}
	}
	return largestArea
}

func main() {

	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}
