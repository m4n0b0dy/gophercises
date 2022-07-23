package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A description")
	timeSeconds := flag.Int("timer", 5, "A description")
	timer1 := time.NewTimer(time.Duration(*timeSeconds) * time.Second) // after time expires, will send a message to channel C
	// ticker time will send every 5 seconds

	f, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatal(err)
	}
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	f.Close()

	if err != nil {
		log.Fatal(err)
	}

	allProblems := parseLines(data)
	allProblems.shuffle()

	score := 0

	for i, p := range allProblems {
		fmt.Printf("Problem %d %s: ", i, p.question)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer1.C:
			fmt.Println("Timer has completed")
			return
		case answer := <-answerCh:
			if answer == p.answer {
				score++
			}
		}
	}
	fmt.Printf("You got %v answers right\n", score)

}

func parseLines(lines [][]string) problems {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

// makes code consistency easier!
type problem struct {
	question string
	answer   string
}

type problems []problem

func (p problems) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range p {
		j := r.Intn(i + 1)
		p[i], p[j] = p[j], p[i]
	}
}
