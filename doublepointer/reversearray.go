package doublepointer

func reverseArray(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	begin, end := 0, len(arr)-1
	for begin < end {
		arr[begin], arr[end] = arr[end], arr[begin]
		begin++
		end--
	}
	return arr
}
