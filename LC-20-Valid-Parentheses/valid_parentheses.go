func isValid(s string) bool {
    
    brackets := []byte{}

    for _, val := range s {
        length := len(brackets)
        lastElem := length - 1 

        switch byte(val){
            case ']':
                 if length == 0 || brackets[lastElem] != '['{
                     return false
                 }
                 brackets = brackets[:lastElem]
            case '}':
                if length == 0 || brackets[lastElem] != '{'{
                     return false
                 }
                 brackets = brackets[:lastElem]
            case ')':
                if length == 0 || brackets[lastElem] != '('{
                    return false
                }
                brackets = brackets[:lastElem]
            default:
                brackets = append(brackets, byte(val))
        }

    }

    return len(brackets) == 0
    
}
