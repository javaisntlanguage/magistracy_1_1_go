package main

import "fmt"

//1.1
func main() {
	fmt.Println(lab1_1([]int{1, 2, 3, 4, 5}))
}

func lab1_1(a []int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}
