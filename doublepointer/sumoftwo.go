package doublepointer

func sumOfTwoNumbers(numbers []int, target int) []int {
	result := make([]int, 2)
	if len(numbers) == 0 {
		return result
	}
	begin, end := 0, len(numbers)-1
	for begin < end {
		if sum := numbers[begin] + numbers[end]; sum == target {
			result[0] = numbers[begin]
			result[1] = numbers[end]
			return result
		} else if sum > target {
			end--
		} else if sum < target {
			begin++
		}
	}
	return result
}
