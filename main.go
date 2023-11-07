package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type DocCounts struct {
	lineCount, wordsCount, vowelsCount, puncuationsCount int
}

func Counts(data string, channal chan DocCounts) {

	DocCounts := DocCounts{}
	fmt.Println(data)

	for _, words := range data {
		switch {
		case words == 13:
			DocCounts.lineCount++
		case words == 32 || words == 10:
			DocCounts.wordsCount++
		case words == 65 || words == 69 || words == 73 || words == 79 || words == 85 || words == 97 || words == 101 || words == 105 || words == 111 || words == 117:
			DocCounts.vowelsCount++
		case (words >= 33 && words <= 47) || (words >= 58 && words <= 64) || (words >= 91 && words <= 96) || (words >= 123 && words <= 126):
			DocCounts.puncuationsCount++
		}

	}

	channal <- DocCounts
}
func main() {
	start := time.Now()
	channal := make(chan DocCounts)

	content, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileData := string(content)

	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	chunk := len(fileData) / a
	startIndex := 0
	endIndex := chunk
	for iterations := 0; iterations < a; iterations++ {
		go Counts(fileData[startIndex:endIndex], channal)
		startIndex = endIndex
		endIndex += chunk
		counts := <-channal
		fmt.Printf("number of lines of chunk %d: %d \n", iterations+1, counts.lineCount)
		fmt.Printf("number of words of chunk %d: %d \n", iterations+1, counts.wordsCount)
		fmt.Printf("number of vowels of chunk %d: %d \n", iterations+1, counts.vowelsCount)
		fmt.Printf("number of puncuations of chunk %d: %d \n", iterations+1, counts.puncuationsCount)

	}

	fmt.Println("Run Time:", time.Since(start))

}
