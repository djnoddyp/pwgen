package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	WORDS_FILE = "/usr/share/dict/british-english"
)

func main() {
	words := readDict()
	filtered := filterByLength(words)
	filtered = filterWhitespace(filtered)


	fmt.Println(getWords(filtered, 4))
}

func readDict() []string {
	content, err := ioutil.ReadFile(WORDS_FILE)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(content[:]), "\n")
}

func filterByLength(data []string) []string {
	lengthPredicate := func(word string) bool {
		return len(word) >= 5 && len(word) <= 8; 
	}
	return Filter(data, lengthPredicate)
}


func getWords(data []string, num int) (words []string) {
	for num > 0 {
		i := RandomNumber(len(data))
		words = append(words, data[i])
		num--
	}
	return
}

func filterWhitespace(data []string) []string {
	words := make([]string, len(data))
	copy(words, data)
	lengthPredicate := func(word string) bool {
		return !strings.Contains(word, "'"); 
	}
	return Filter(words,  lengthPredicate)
}

