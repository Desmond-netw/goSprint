package main

func ContainDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	// this sore the number in hashmap as key and struct as empty to use zero memeory

	// iterate to find the number
	for _, num := range nums {

		// put number in keys to check if it esit
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{}
	}

	return false
}
