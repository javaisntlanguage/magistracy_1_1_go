package main

import
(
	"fmt"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func minEl(arr []int, minValCh chan int) {
	min := MaxInt
	for _, val := range arr {
		if val < min {
			min = val
		}
	}

	minValCh <- min
}

func maxEl(arr []int, maxValCh chan int) {
	max := MinInt
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	maxValCh <- max
}

func main() {
	arr := []int{-67, 2, -100, 10, 5, 3}
	minValCh := make(chan int)
	maxValCh := make(chan int)

	var minVal int
	var maxVal int

	go minEl(arr, minValCh)
	go maxEl(arr, maxValCh)

	for i := 0; i < 2; i++ {
		select {
		case minVal = <-minValCh:
		case maxVal = <-maxValCh:
		}
	}

	println(fmt.Sprintf("minVal %d", minVal))
	println(fmt.Sprintf("maxVal %d", maxVal))
}