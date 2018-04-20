package Arrays

// Takes a list of numbers and sums them together
func Sum(numbers []int) int {
	total := 0
	for _, i := range numbers {
		total += i
	}
	return total
}

// Takes multiple lists and sums them
func SumAll(numsToSum ...[]int) []int {
	var sums []int
	for _, nums := range numsToSum {
		sums = append(sums, Sum(nums))
	}
	return sums
}

func SumAllTails(numsToSum ...[]int) []int {
	var sums []int
	for _, nums := range numsToSum {
		if len(nums) != 0 {
			sums = append(sums, Sum(nums[1:]))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
