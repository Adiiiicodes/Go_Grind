package main
import ("fmt" 
		"algo/algo" 
		"time")

func main() {
	arr := []int{100 , 2 , 1 , 4 , 81 , 6 , 59 , 7}
	start := time.Now()
	sortedArray := algo.Mergesort(arr)
	totaltime := time.Since(start)
	fmt.Printf("time consumed: %v \n", totaltime)
	fmt.Println("Sorted Array is --> ", sortedArray)
}
