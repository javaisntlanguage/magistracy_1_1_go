package main

import
(
	"fmt"
	"sort"
)

func main(){
	GetMedInArr()
}

func GetMedInArr(){

	list := []int{5,76,3,6,7,3,36,7,82,5,7}
	// сортировка
	sort.Ints(list)

	if((len(list)%2) == 1){
		ind := int(len(list)%2)
		med := list[ind]
		fmt.Println("Медиана: ", med, ". индекс: ", int(len(list)%2) + 1)
	} else {
		first := len(list)/2
		second := len(list)/2-1
		med := (list[first] + list[second])/2;
		fmt.Println("Медиана :", med, ". индекс: ", first+1, " и ", second+1)
	}
}