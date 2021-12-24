package main

import
(
	"io/ioutil"
	"strconv"
)

func findMaxSizeFile(file_path string, maxSize int64, maxFileName string) (int64, string) {
	files, err := ioutil.ReadDir(file_path)
	if err != nil {
		return maxSize, maxFileName
	}

	for _, f := range files {
		fSize := f.Size()
		name := f.Name()
		if fSize > maxSize {
			maxSize = fSize
			maxFileName = name
		}

		maxSize, maxFileName = findMaxSizeFile(file_path+f.Name()+"/", maxSize, maxFileName)
	}

	return maxSize, maxFileName
}

func main() {
	file_path := "./"
	maxSize, maxFileName := findMaxSizeFile(file_path, -1, "./")
	println(strconv.FormatInt(maxSize, 10) + " " + maxFileName)

	println()
}