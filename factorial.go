package main

import "fmt"

func factorial(num int) int {
	if num == 1 || num == 0 {
		return num
	}
	return num * factorial(num-1)
}
func main() {

	fmt.Printf("Enter size of your array: ")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanf("%d", &arr[i])
		fmt.Println(factorial(arr[i]))
	}

}
