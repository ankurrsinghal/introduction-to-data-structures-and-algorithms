package lecture2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSpanOfStocksBetter(t *testing.T) {

	array := []int{6, 4, 1, 3, 1, 4, 5}
	spansArray := spanOfStockBetter(array)
	expectedArray := []int{1, 1, 1, 2, 1, 5, 6}

	fmt.Println(spansArray)

	if !reflect.DeepEqual(spansArray, expectedArray) {
		t.Error("Not Equal")
	}
}

func TestSpanOfStocks(t *testing.T) {
	array := []int{6, 4, 1, 3, 1, 4, 5}
	spansArray := spanOfStock(array)
	expectedArray := []int{1, 1, 1, 2, 1, 5, 6}

	fmt.Println(spansArray)

	if !reflect.DeepEqual(spansArray, expectedArray) {
		t.Error("Not Equal")
	}
}

func TestGrowableStackGrowthStrategy(t *testing.T) {
	var stack Stack = createGrowableStackGrowthStrategy()

	if !stack.isEmpty() {
		t.Error("isEmpty() Failed")
	}

	stack.push(1)

	if stack.size() != 1 {
		t.Error("Push() Failed")
	}

	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(1)

	if stack.top() != 1 {
		t.Error("Top() Failed")
	}

	stack.pop()
	if stack.top() != 4 {
		t.Error("Pop() Failed")
	}
}

func TestGrowableStackTightStrategy(t *testing.T) {
	var stack Stack = createGrowableStackTightStrategy()

	if !stack.isEmpty() {
		t.Error("isEmpty() Failed")
	}

	stack.push(1)

	if stack.size() != 1 {
		t.Error("Push() Failed")
	}

	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(1)

	if stack.top() != 1 {
		t.Error("Top() Failed")
	}

	stack.pop()
	if stack.top() != 4 {
		t.Error("Pop() Failed")
	}
}

func TestStack(t *testing.T) {
	var stack Stack = createStack(10)

	if !stack.isEmpty() {
		t.Error("isEmpty() Failed")
	}

	stack.push(1)

	if stack.size() != 1 {
		t.Error("Push() Failed")
	}

	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.push(1)

	if stack.top() != 1 {
		t.Error("Top() Failed")
	}

	stack.pop()
	if stack.top() != 4 {
		t.Error("Pop() Failed")
	}
}
