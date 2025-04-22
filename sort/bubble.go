package sort

func BubbleSort(arr []float64) {
	swap := true
	end := len(arr)
	for swap {
		swap = false
		for i := 1; i < end; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swap = true
			}
		}
		end--
	}
}
