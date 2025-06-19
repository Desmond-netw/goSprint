package main


func TwoSum(numbs []int, target int) []int {
	// make a space and map number index to their indices
	myMap := make(map[int]int)

	// loop over rang of array
	for index, num := range numbs {
		// calculate with formular
		complement := target - num
		if idx, ok := myMap[complement]; ok {
			return []int{idx, index}
		}
		myMap[num] = index

	}
	return nil 
}