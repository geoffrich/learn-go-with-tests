package sum

// []int -> slice, does not have fixed size
// [3]int -> array, has fixed size (size encoded in type)
func Sum(numbers []int) int {
	sum := 0
	// range iterates over an array, returning index and value
	// we are ignoring index with the blank identifier
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// variadic functions -- takes arbitrary number of []ints as an argument
	var sums []int

	for _, numbers := range numbersToSum {
		// create new slice with additional item (since slices have a fixed size)
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// slices can be sliced!
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
