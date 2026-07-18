package algo1
//import ("fmt" "time" )
		

func Bubblesort(arr []int) []int {
	//start := time.Now()
	size := len(arr) -1
	for i:=0; i<size; i++ {
		if arr[i] > arr[i+1] {
			arr[i] , arr[i+1] = arr[i+1] , arr[i]
		}
	} 

	if  !sorted(arr) {
		Bubblesort(arr)
	}

//	totaltime:= time.Since(start)
	//fmt.Printf("time consumed: %v \n", totaltime)
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