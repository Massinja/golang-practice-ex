func majorityElement(nums []int) int {
    m := make(map[int]int)
        
    for _, val := range nums {
        m[val]++
    }
    majSize := (len(nums)) / 2
        
    for val, k := range m {
        if k > majSize {
            return val
        }
    }

    return 0

}
