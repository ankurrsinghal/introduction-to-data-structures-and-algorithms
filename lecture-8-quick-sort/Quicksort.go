package lecture8

func partition(A []int, p int, r int) int {
	pivot := A[r]

	for p <= r {
		if A[p] >= pivot && A[r] <= pivot {
			A[p], A[r] = A[r], A[p]
		}
		if A[p] <= pivot {
			p++
		}
		if A[r] >= pivot {
			r--
		}
	}

	return r
}

// QuickSort ...
func QuickSort(A []int, p int, r int) {
	if p < r {
		q := partition(A, p, r)
		QuickSort(A, p, q)
		QuickSort(A, q+1, r)
	}
}
