package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

type result struct {
	string
	int
}

func fileRead(fi string) map[string]int {
	data, err := ioutil.ReadFile(fi)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(data))
	text := string(data)

	fields := strings.FieldsFunc(text, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})
	words := make(map[string]int)
	for _, field := range fields {
		words[strings.ToLower(field)]++
	}
	// fmt.Println(words)
	return words
}

func folderRead(fo string) []string {
	var arrDir []string
	files, err := ioutil.ReadDir(fo)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if string(file.Name())[len(string(file.Name()))-3:] == "txt" {
			arrDir = append(arrDir, string(file.Name()))
		}
	}
	// fmt.Println(arrDir)
	return arrDir
}

func main() {

	resultMap := make(map[string]int)
	resultChannel := make(chan result)

	for _, file := range folderRead(".") {
		go func(w string) {
			resultChannel <- result{w, fileRead(w)[w]}
		}(file)
	}

	for i := 0; i < len(folderRead(".")); i++ {
		result := <-resultChannel
		resultMap[result.string] = result.int
	}
	fmt.Println(resultMap)

}
