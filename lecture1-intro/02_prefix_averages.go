package lecture1

func prefixAveragesBetter(array []int) []int {
	result := make([]int, len(array))
	var sum int
	for i := 0; i < len(array); i++ {
		sum += array[i]
		result[i] = sum / (i + 1)
	}

	return result
}

func prefixAverages(array []int) []int {
	result := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += array[j]
		}
		result[i] = sum / (i + 1)
	}

	return result
}
