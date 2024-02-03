package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {


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