package main

func isPalindrome(s string) bool {
	length := len(s)

	if length < 1 {
		return false
	}

	left := 0
	right := length - 1
	str := []byte(s)

	for left < right {
		for !isAlphanumeric(rune(s[left])) && left < length-1 { //"   ," length 4 left = 3 ","
			left++
		}
		if left == length-1 {
			return true
		}
		for !isAlphanumeric(rune(s[right])) && right > 0 { // right = 0 " "
			right--
		}

		if lowerCase(str[left]) != lowerCase(str[right]) {
			return false
		} else {
			left++
			right--
		}

	}

	return true

}

func lowerCase(c byte) byte {
	if rune(c) >= 'A' && rune(c) <= 'Z' {
		c = byte(rune(c) + 32)
	}
	return c

}

func isAlphanumeric(c rune) bool {
	return (c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}
