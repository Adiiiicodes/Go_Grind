package main
import ("fmt" 
		"algo/algo")

func main() {
	arr := []int{5,8,1,6,7,3,4,2}
	sortedArray := algo.Mergesort(arr)
	fmt.Println("Sorted Array is --> ", sortedArray)
}
