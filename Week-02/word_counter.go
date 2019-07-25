package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"unicode"
)

const MAX_POOL = 2

func getNextWord(file *os.File) (string, error) {
	var buff = make([]byte, 1) // TODO: Optimize buffer size to improve reading performance
	var stringBuilder []rune

	for {
		_, err := file.Read(buff)

		if err == io.EOF && len(stringBuilder) > 0 {
			return string(stringBuilder), nil
		}

		if err != nil {
			return "", err
		}

		switch c := rune(buff[0]); {
		case unicode.IsLetter(c):
			fallthrough
		case unicode.IsDigit(c):
			stringBuilder = append(stringBuilder, c)
		case c == ' ':
			if len(stringBuilder) > 0 {
				return string(stringBuilder), nil
			}
		}
	}
}

func wordCounter(wg *sync.WaitGroup, files chan string, words chan string) {
	// Do until files chanel close
	for file := range files {
		// Not check error here since if an error occurred that will return when call getNextWord
		f, _ := os.Open(file)
		defer f.Close()

		for {
			word, err := getNextWord(f)

			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			words <- word
		}
	}

	wg.Done()
}

func findFiles(dirPath string, files chan string) {
	err := filepath.Walk(dirPath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.Mode().IsRegular() {
				files <- path
			}

			return nil
		})

	if err != nil {
		log.Fatal(err)
	}

	close(files) // Close files chanel
}

func resultCounter(words chan string, wordMap map[string]uint64) {
	for word := range words {
		wordMap[word]++
	}
}

func main() {
	var i uint32
	var files = make(chan string)
	var words = make(chan string)
	var path = "./word_counter_data_test/"
	var wordCounterWG sync.WaitGroup
	var wordMap = make(map[string]uint64)

	for i = 0; i < MAX_POOL; i++ {
		wordCounterWG.Add(1)
		go wordCounter(&wordCounterWG, files, words)
	}

	go resultCounter(words, wordMap)

	findFiles(path, files)

	wordCounterWG.Wait()

	fmt.Println(wordMap)
}
