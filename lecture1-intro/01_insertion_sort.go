package lecture1

func insertionSort(array []int) []int {
	for j := 1; j < len(array); j++ {
		key := array[j]

		var i int
		for i = j - 1; i >= 0 && array[i] > key; i-- {
			array[i+1] = array[i]
		}

		array[i+1] = key
	}

	return array
}
