package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {

	lines:= parseCsv()
	problems:= parseLines(lines)
	timer:=startTimer()
	quiz(problems, timer)

}


func parseCsv() [][] string {
	filename:= flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*filename)
	if err!= nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *filename))
	}

	r:= csv.NewReader(file)

	lines, err:= r.ReadAll()
	if err!= nil {
		exit("Failed to parse the CSV file.")
	}

	return lines

}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {

	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret

}

func startTimer() *time.Timer {

	flag.Parse()
	timeLimit:= flag.Int("limit", 30, "the time limit for the quiz in seconds")

	timer:= time.NewTimer(time.Duration(*timeLimit)* time.Second)
	return timer
}

func quiz(problem []problem ,timer *time.Timer) {
	correct:= 0

	problemloop:
	for i, p := range problem {
		fmt.Printf("Problem #%d: %s =", i+1, p.q)
		answerCh:= make(chan string)
		go func ()  {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				correct ++
			}

		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problem))

}


func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}