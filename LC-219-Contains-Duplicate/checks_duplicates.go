func containsNearbyDuplicate(nums []int, k int) bool {
     m := make(map[int]int)
     for i, val := range nums {
         if _, ok := m[val]; ok {
             if (i - m[val]) <= k{
                 return true
             }
         }
         m[val]=i
     }
     return false
}
