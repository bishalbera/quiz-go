package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	lines:= parseCsv()
	problems:= parseLines(lines)
	quiz(problems)

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

func quiz(problem []problem) {
	correct:= 0

	for i,p := range problem {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}

	}
	fmt.Printf("You scored %d out of %d. \n", correct, len(problem))
}


func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}