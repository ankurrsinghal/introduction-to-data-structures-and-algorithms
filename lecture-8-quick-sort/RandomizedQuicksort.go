package lecture8

import (
	"fmt"
	"math/rand"
)

func randomizedPartition(A []int, p int, r int) int {
	i := rand.Int31n(int32(r-p+1)) + int32(p)
	A[i], A[r] = A[r], A[i]
	return partition(A, p, r)
}

// RandomizedQuickSort ...
func RandomizedQuickSort(A []int, p int, r int) {
	fmt.Println(A)
	if p < r {
		q := randomizedPartition(A, p, r)
		RandomizedQuickSort(A, p, q)
		RandomizedQuickSort(A, q+1, r)
	}
}
