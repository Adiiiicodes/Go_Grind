package algo1
		

func Bubblesort(arr []int) []int {
	size := len(arr) -1
	for i:=0; i<size; i++ {
		if arr[i] > arr[i+1] {
			arr[i] , arr[i+1] = arr[i+1] , arr[i]
		}
	} 

	if  !sorted(arr) {
		Bubblesort(arr)
	}

	return arr
	
}

func sorted(arr []int) bool {
	size := len(arr) -1
	for i:=0; i<size; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	} 
	return true
}