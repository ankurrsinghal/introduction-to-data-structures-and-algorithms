package lecture8

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	array := []int{3, 2, 9, 5, 1, 1}
	QuickSort(array, 0, len(array)-1)
	expected := []int{1, 1, 2, 3, 5, 9}

	if !reflect.DeepEqual(array, expected) {
		t.Error("Sorting failed")
	}
}

func TestRandomizedQuickSort(t *testing.T) {
	array := []int{3, 2, 9, 5, 1, 1}
	RandomizedQuickSort(array, 0, len(array)-1)
	expected := []int{1, 1, 2, 3, 5, 9}

	if !reflect.DeepEqual(array, expected) {
		t.Error("Sorting failed")
	}
}
