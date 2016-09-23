package sortingBenchmark

func mergeSort(values []int) []int {
	if len(values) < 2 {
		return values
	}
	mid := len(values) / 2
	left := mergeSort(values[:mid])
	right := mergeSort(values[mid:])

	return merge(left, right)
}

func parallelMergeSort(data []int, r chan []int) {
	if len(data) == 1 {
		r <- data
		return
	}

	leftChan := make(chan []int)
	rightChan := make(chan []int)
	middle := len(data) / 2

	go parallelMergeSort(data[:middle], leftChan)
	go parallelMergeSort(data[middle:], rightChan)

	ldata := <-leftChan
	rdata := <-rightChan

	close(leftChan)
	close(rightChan)
	r <- merge(ldata, rdata)
	return
}

func merge(left, right []int) []int {
	var i, j int
	result := make([]int, len(left)+len(right))

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result[i+j] = left[i]
			i++
		} else {
			result[i+j] = right[j]
			j++
		}
	}

	// put the rest of left and right side in result
	for i < len(left) {
		result[i+j] = left[i]
		i++
	}

	for j < len(right) {
		result[i+j] = right[j]
		j++
	}

	return result
}
