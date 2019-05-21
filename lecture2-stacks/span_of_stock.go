package lecture2

// spanOfStock :- Using stack
func spanOfStockBetter(stocks []int) []int {
	result := make([]int, len(stocks))
	stack := createStack(len(stocks))

	for i, price := range stocks {
		h := -1
		for !stack.isEmpty() {
			top, ok := stack.top().(int)
			if ok {
				topPrice := stocks[top]
				if price < topPrice {
					h = top
					break
				} else {
					stack.pop()
				}
			} else {
				panic("Not an integer")
			}
		}

		result[i] = i - h
		stack.push(i)
	}

	return result
}

// spanOfStock :- Without using stack
func spanOfStock(stocks []int) []int {
	result := make([]int, len(stocks))

	for i, price := range stocks {
		span := 1
		for j := i - 1; j >= 0 && stocks[j] <= price; j-- {
			span++
		}
		result[i] = span
	}

	return result
}
