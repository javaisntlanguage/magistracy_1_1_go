package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)
//2.1
func main() {

	path,_ := os.Getwd()
	names:= GetImgFileNames(path + "\\images")
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	
}

func GetImgFileNames(path string) []string{
	var names []string
	files, err := ioutil.ReadDir(path)

	if err != nil {
        log.Fatal(err)
    }
 
    for _, f := range files {
		pattern := "\\.(jpg|png|bmp|gif|ico|jpeg)$"
		matched,_ := regexp.MatchString(pattern, f.Name())
		if matched{
            names = append(names, f.Name())
		}
    }

	return names
}
