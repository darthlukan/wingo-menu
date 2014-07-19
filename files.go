package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const filesDir string = "/usr/share/applications/"

func main() {
	files, err := ioutil.ReadDir(filesDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Printf("Found: %s\n", file.Name())
		filePath := fmt.Sprintf("%s%s", filesDir, file.Name())
		f, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			lineArray := strings.SplitN(scanner.Text(), "=", 2)
			if len(lineArray) > 1 {
				fmt.Printf("Entry: %s: %s\n", lineArray[0], lineArray[1])
			}
		}
		if err = scanner.Err(); err != nil {
			fmt.Printf("error: %s\n", err)
		}
	}
}
