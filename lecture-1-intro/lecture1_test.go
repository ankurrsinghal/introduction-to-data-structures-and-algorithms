package lecture1

import (
	"reflect"
	"testing"
)

func TestPrefixesAverageBetter(t *testing.T) {
	array := []int{2, 4, 6}
	avgArray := prefixAverages(array)
	expectedArray := []int{2, 3, 4}

	if !reflect.DeepEqual(avgArray, expectedArray) {
		t.Error("Not Equal")
	}
}
func TestPrefixesAverage(t *testing.T) {
	array := []int{2, 4, 6}
	avgArray := prefixAverages(array)
	expectedArray := []int{2, 3, 4}

	if !reflect.DeepEqual(avgArray, expectedArray) {
		t.Error("Not Equal")
	}
}

func TestInsertionSort(t *testing.T) {
	array := []int{2, 5, 3, 1, 9, 6, 7}
	sortedArray := insertionSort(array)
	expectedArray := []int{1, 2, 3, 5, 6, 7, 9}

	if !reflect.DeepEqual(sortedArray, expectedArray) {
		t.Error("Not Equal")
	}
}
