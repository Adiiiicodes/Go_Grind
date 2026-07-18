package algo


func Insertionsort(arr []int) []int {
	size := len(arr)
	for i := 1; i <size; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
	return arr
}