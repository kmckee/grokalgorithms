package grokalgorithms

// BinaryChop finds an item within an array by checking the median
// item in the slice against the desired item to see which half of
// the slice to keep and which half to throw out.  log2 n is the worst
// case.
func BinaryChop(items []int, lookFor int) int {
	low := 0
	high := len(items) - 1
	for low <= high {
		mid := int((low + high) / 2)
		current := items[mid]
		if lookFor == current {
			return mid
		} else if lookFor < current {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
