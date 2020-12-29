package gen

import (
	"io/ioutil"
	"log"
	"strings"
)

func Generate() string {
	words := readDict()
	filtered := filterByLength(words, 5, 8)
	filtered = filterWhitespace(filtered)
	return strings.Join(getWords(filtered, 4), "")
}

func readDict() []string {
	content, err := ioutil.ReadFile(WordsFile)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(content[:]), "\n")
}

func filterByLength(data []string, min, max int) []string {
	lengthPredicate := func(word string) bool {
		return len(word) >= min && len(word) <= max
	}
	return Filter(data, lengthPredicate)
}

func getWords(data []string, num int) (words []string) {
	for num > 0 {
		i := RandomNumber(len(data))
		words = append(words, strings.Title(data[i]))
		num--
	}
	return
}

func filterWhitespace(data []string) []string {
	words := make([]string, len(data))
	copy(words, data)
	lengthPredicate := func(word string) bool {
		return !strings.Contains(word, "'")
	}
	return Filter(words, lengthPredicate)
}
