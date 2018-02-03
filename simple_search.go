package grokalgorithms

// SimpleSearch performs a simple beginning to end search for an
// item in the items array.  n steps is the worst case,
// which would happen if the item we're looking for is last.
func SimpleSearch(items []string, lookFor string) int {
	for index, item := range items {
		if item == lookFor {
			return index
		}
	}
	return -1
}
