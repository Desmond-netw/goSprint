package main


func twoSum(numbers []int, target int) []int {
	// initialize result
	result := []int{}

	// loop through arry of number
	for i := 0; i < len(numbers); i++ {
		for j := 0; i < len(numbers); j++ {
			// condition
			if numbers[i] + numbers[j] == target && i != j {
				result = []int{i, j}
				return result
			}
		}
	}
	return nil
}
