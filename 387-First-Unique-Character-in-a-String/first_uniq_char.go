func firstUniqChar(s string) int {
    m := make(map[byte]int)

    for i, val := range s {
        if _, ok := m[byte(val)]; ok {
            m[byte(val)] = -1             
        }else{
           m[byte(val)] = i 
        }
              
    }
    lowestIndex := 1000000

    for _, value := range m {
        if value != -1 && value < lowestIndex {
            lowestIndex = value
        }

    }
    if lowestIndex == 1000000 {
        return -1
    }
    return lowestIndex
    
}
