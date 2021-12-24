package main

import (
	"bufio"
	"fmt"
	"os"
	_ "os"
	"sort"
	"strconv"
	"strings"
	"math/rand"
	"io/ioutil"
)

type Person struct {
	age int
	id int
	lastName, firstName, middleName string
	email []string
	url string
}

type Persons []Person

func NewPerson(lastName string, firstName string, middleName string, email []string, age int, url string) Person{
	return Person{ age: age, lastName: lastName, firstName: firstName, middleName: middleName, email: email,
	url: url, id: rand.Intn(100000)}
}

// добавить адрес
func (person Person)AddEMail(email string) {
	person.email = append(person.email, email)
}

// удалить email по значению
func (person Person)RemoveEMail(email string) {
	//получаем индекс
	i := sort.SearchStrings(person.email, email)
	//смещаем массив
	person.email = append(person.email[:i], person.email[i+1:]...)
}

// очищаем массив url
func (person Person)CrealEMail() {
	person.email = []string{}
}

// поиск по фамилии
func (persons Persons) FindByLastName(s string)  {
	fmt.Println("Поиск по фамилии \"", s,"\":")
	for _, i := range persons {
		if (strings.Contains(i.lastName, s)){
			fmt.Println(i.lastName)
		}
	}
}
// поиск по соц.сетям
func (persons Persons) FindByUrl(s string)  {
	fmt.Println("Поиск по соц. сетям\"", s,"\":")
	for _, i := range persons {
		if (strings.Contains(i.url, s)){
			fmt.Println(i.lastName, " ", i.firstName, " ",  i.middleName)
		}
	}
}

// запись в файл
func (person Person)SaveToFile() {
	text := person.lastName + " " + person.firstName + " " + person.middleName
	filename := "files/" + person.lastName + strconv.Itoa(person.id) + ".txt"
	ioutil.WriteFile(filename, []byte(text), 0666)
	file, err := os.OpenFile(filename, os.O_CREATE, 0666)
	if err != nil{
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}
	file.WriteString(text)
	//вызывается после завершения работы метода
	defer file.Close()
	fmt.Println("Успешно записано в файл.")
}

// загрузка из файла
func (person *Person)ReadFromFile(filename string)  {

	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		person.lastName = words[0]
		person.firstName = words[1]
		person.middleName = words[2]
	}
	defer file.Close()
}