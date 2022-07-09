package reverse_string

func ReverseString(input string) (output string) {
	tempRune := []rune(input)
	tempRuneLen := len(tempRune)
	newRune := make([]rune, tempRuneLen)

	// iterate tempRune backwards and
	// assign the last symbols of tempRune to
	// initial symbols of empty newRune
	for i := tempRuneLen - 1; i >= 0; i-- {
		newRune[tempRuneLen-i-1] = tempRune[i]
	}
	output = string(newRune)

	return output
}
