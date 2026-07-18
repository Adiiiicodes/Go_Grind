package main
import ("fmt"
		"algo1/algo1"
		"time")

func main() {
	arr := []int{5,2,4,6,2,7,8 , 7}
	start := time.Now()
	sortedArray := algo1.Bubblesort(arr)
	totaltime := time.Since(start)
	fmt.Printf("time consumed: %v \n", totaltime)
	fmt.Println("Sorted Array is --> ", sortedArray)
}

