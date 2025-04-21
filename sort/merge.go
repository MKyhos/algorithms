package sort

func merge(left, right []int) []int {
	out := make([]int, len(left)+len(right))
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			out[k] = left[j]
			i++
		} else {
			out[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		out[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		out[k] = right[j]
		j++
		j++
	}
	return out
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return merge(left, right)
}
