package main
import ("fmt"
		"algo1/algo1" )

func main() {
	arr := []int{100 , 2 , 1 , 4 , 81 , 6 , 59 , 7}
	sortedArray := algo1.Bubblesort(arr)
	fmt.Println("Sorted Array is --> ", sortedArray)
}

