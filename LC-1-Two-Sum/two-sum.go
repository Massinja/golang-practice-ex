package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, val := range nums {
	 	if k, ok := m[target - val]; ok {
			return []int{k,i}
		}
		m[val] = i
		
	}
	return []int{0,0}
}

