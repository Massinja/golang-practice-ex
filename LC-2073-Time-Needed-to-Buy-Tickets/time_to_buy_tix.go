func timeRequiredToBuy(tickets []int, k int) int {
   
    total := 0
        
    for i, val := range tickets {
        if i > k{
            if val < tickets[k]{
                total += val
            }else{
                total += tickets[k] - 1
            }
        }else{
            if val < tickets[k]{
                total += val
            }else{
                total += tickets[k]
            }
        }
    }
    
    return total
}
