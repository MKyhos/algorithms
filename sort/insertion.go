package sort

func InsertionSort(arr []int) {
	for i := range arr {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			tmp := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = tmp
			j--
		}
	}
}
