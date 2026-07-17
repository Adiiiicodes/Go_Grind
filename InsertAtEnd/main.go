package main
import "fmt"

func main () {
	var arr = [5]int{9 , 8 , 7 , 5 , 2 }
	fmt.Printf("Array before shifting: %d\n", arr)

	size := len(arr) -1
	value := 96

	for i:= 0; i<size; i++ {
		arr[i] = arr[i+1]
	}

	arr[size] = value
	fmt.Printf("Array after shifting: %d\n", arr)
}
