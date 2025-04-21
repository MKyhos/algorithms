package sort

import "errors"

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	i := lo
	for j := lo; j < hi; j++ {
		if arr[j] <= pivot {
			tmp := arr[j]
			arr[j] = arr[i]
			arr[i] = tmp
			i++
		}
	}
	tmp := arr[i]
	arr[i] = arr[hi]
	arr[hi] = tmp
	return i
}

func QuickSort(arr []int, lo, hi int) error {
	if lo > hi || lo < 0 {
		return errors.New("bounds not aligned")
	}

	p := partition(arr, lo, hi)
	QuickSort(arr, lo, p-1)
	QuickSort(arr, p+1, hi)
	return nil
}
