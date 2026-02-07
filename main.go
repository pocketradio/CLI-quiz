package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	csvfilename := flag.String("csv", "problems.csv", "a csv file for problems")
	flag.Parse()

	timelimit := flag.Int("limit", 30, "time limit for quiz in seconds")

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
	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)
	// <- timer.C // newtimer returns a timer struct that has a channel 'C' by def
	// this blocks current goroutine (main) until the timer fires.

	correct := 0

problemloop:
	for i, p := range problems {
		fmt.Printf("Question No. %d : %s = \n", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s\n", &ans)
			answerCh <- ans
		}()

		select {
		case <-timer.C: // conceptually it receives a timestamp repr when the timer fired. but we dont need that here since we only need to know when it got over ( this is confirmed by the channel recieving something ie.) ; so the value is intentionally discarded
			fmt.Printf("\n you scored %d out of %d.\n", correct, len(problems))
			break problemloop

		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	//needed in case user answers all q within timelimit
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
