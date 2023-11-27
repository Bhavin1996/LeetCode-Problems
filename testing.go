package leetcode

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main_file() {
	dir := "/home/raiden/Desktop/FILE_IO/"
	fileDir, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Issue with the direcetory while accesing")
	}
	re, err := regexp.Compile(`^(25[0-5]|2[0-4][0-9]|[0-1][0-9][0-9]).{3}(25[0-5]|2[0-4][0-9]|[0-1][0-9][0-9])`)

	for _, fileName := range fileDir {
		filePath := filepath.Join(dir, fileName.Name())
		file, err := os.ReadFile(filePath)

		for _, data := range file {
			words := strings.Fields(string(data))

			for _, word := range words {
				if re.MatchString(word) && err != nil {
					fmt.Println("Valid IP address ", word)
				}
			}

		}

	}

}
