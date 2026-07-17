package main
import ("fmt" 
		"algo/algo")

func main() {
	arr := []int{5,8,1,6,7,3,4,2}
	sortedArray := algo.Mergesort(arr)
	fmt.Println("Sorted Array is --> ", sortedArray)
}

/*func mergesort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr)/2
	left := mergesort(arr[:mid])
	right := mergesort(arr[mid:])
	return merge(left, right)
}

func merge (left , right []int) []int {
	i , j := 0 , 0 
	result := []int{}
	for i< len(left) && j<len(right) {
		if left[i]< right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result , right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}*/