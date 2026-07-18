package main
import ( "fmt"
		"time" 
		"algo/algo")
	

func main() {
	
	arr := []int{5,2,4,6,2,7,8 , 7}
	start := time.Now()
	sortedarray := algo.Insertionsort(arr)
	end := time.Since(start)
	fmt.Println("Sorted Array is --> ", sortedarray)
	fmt.Println("Time taken: ", end)
}


/*func Insertionsort(arr []int) []int {
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
}*/