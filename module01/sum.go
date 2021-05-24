package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func RecursiveSum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + RecursiveSum(numbers[1:])
}
