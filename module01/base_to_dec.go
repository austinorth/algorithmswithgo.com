package module01

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	const charset = "0123456789ABCDEF"
	res := 0
	multiplier := 1
	for i := len(value) - 1; i >= 0; i-- {
		index := -1
		for j, char := range charset {
			if char == rune(value[i]) {
				index = j
				break
			}
		}
		res += index * multiplier
		multiplier = multiplier * base
	}
	return int(res)
}
