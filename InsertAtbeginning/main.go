package main

import "fmt"

func  main() {
	var arr = [5]int{5 , 15 , 45 , 85 , 92}
	fmt.Println(arr)

	size := len(arr) -1
	value := 10

	for i := size; i>0; i-- {
		arr[i] = arr[i-1]
	}
	fmt.Println("array after shifting: ",arr)

	arr[0] = value
	fmt.Println("array after insertion at beginning: ",arr) 
}