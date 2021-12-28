package main

import (
	"fmt"
	"strings"
)

func main(){
	FindLongestWord()
}


func FindLongestWord()  {	
	text := "Числа используются повсюду в программировании." +
	"Они используются для представления таких вещей, как размеры экрана, "+
	"географическое местоположение, денежные средства и баллы, количество времени, "+
		"передаваемого в видео, расположение игровых аватаров, отображение цветов через присвоение числовых кодов";
	maxLen := 0
	indMaxLen := 0
	words := strings.Fields(text)
	for idx, word := range words {
		if (maxLen < len(word)){
			maxLen = len(word)
			indMaxLen = idx
		}
	}
	fmt.Println("слово: ", words[indMaxLen], " длина: ", maxLen)
}