package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	csvfilename := flag.String("csv", "problems.csv", "a csv file for problems")
	flag.Parse()

	file, err := os.Open(*csvfilename)
	// csvfilename is a pointer to a string. so we need to deref it to get the string itself

	if err != nil {
		exit(fmt.Sprintf("failed to open csv file : %s\n", *csvfilename))
	}

	// now create a csv reader since the file is now open
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse file.")
	}

	problems := parseLines(lines)

	correct := 0
	for i, p := range problems {
		fmt.Printf("Question No. %d : %s = \n", i+1, p.q)
		var ans string
		fmt.Scanf("%s\n", &ans)
		if ans == p.a {
			correct++
		}
	}

	fmt.Printf("you scored %d out of a total %d questions", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	res := make([]problem, len(lines))
	for i, line := range lines {
		res[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]), // trimspace will fix possible incorrect white spaces in the csv file itself. scanf will remove spaces directly from user input.
		}
	}

	return res
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Print(msg)
	os.Exit(1) // exit 1 -> means appln had some error.
}
