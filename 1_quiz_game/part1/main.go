package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A description")
	f, err := os.Open(*csvFilename) // csvfilename is pointer to string so use * to get original
	// we want the actual value
	if err != nil {
		log.Fatal(err)
	}

	// IMPORTANT this might make sense to break up these sections into sub functions so you can properly test

	// want to use a csv reader
	// io.reader is the most common interface used (just like udemy http thing)
	// takes byteslice to read into and returns read
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	f.Close()

	if err != nil {
		log.Fatal(err)
	}
	problems := parseLines(data) //data must be a slice of byte slices
	score := 0
	for i, p := range problems {
		fmt.Printf("Problem %d %s: ", i, p.question) // question attribute of problem struct
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			score++
		}
	}

	fmt.Printf("You got %v answers right\n", score)

}

// takes in lines which is a slice of byte slices
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines)) // make a problem slice (slice filled with problems of length line
	// we pre define length since we know exactly how big it has to be at the start
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]), // small thing but allows trimming of spaces in answer col
		}
	}
	return ret
}

// makes code consistency easier!
type problem struct {
	question string
	answer   string
}
