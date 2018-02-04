package grokalgorithms

// BinaryChop finds an item within an array by checking the median
// item in the slice against the desired item to see which half of
// the slice to keep and which half to throw out.  log2 n is the worst
// case.
func BinaryChop(items []int, lookFor int) int {
	low, high := getRangeFromSlice(items)
	for low <= high {
		mid := getMedianIndex(low, high)
		current := items[mid]
		if lookFor == current {
			return mid
		}

		if lookFor < current {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func getRangeFromSlice(items []int) (low int, high int) {
	low = 0
	high = len(items) - 1
	return
}

func getMedianIndex(min, max int) int {
	return int((min + max) / 2)
}
