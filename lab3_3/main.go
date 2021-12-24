package main

import (
	"fmt"
	_ "math"
	"strconv"
)

func main(){
	var p1 = NewPerson("Воронов", "Андрей", "Витальевич", []string{"mos@mail.ru"}, 22, "site.ru")
	var p2 = NewPerson("Иванов", "Иван", "иванович", []string{"memos@mail.ru", "croco@mail.ru" }, 20, "site2.ru")
	var p3 = NewPerson("Миронова", "Дарья", "Александровна", []string{"123@mail.ru"}, 30, "kolobok.ru")

	var people = Persons{p1, p2, p3}

	people.FindByLastName("он")

	people.FindByUrl("sit")

	p2.SaveToFile()

	filename := "files/" + p2.lastName + strconv.Itoa(p2.id) + ".txt"
	fmt.Println("Содержимое файла : " + filename)
	p3.ReadFromFile(filename)
	fmt.Println(p3.lastName + " " + p3.firstName + " " + p3.middleName)

}