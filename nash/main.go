package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {

	filename := flag.String("filename", "problems.csv", "ex. -filename=problems.csv")
	timeout := flag.Int("timeout", 30, "Amount of time in sec the palyer has to answer a question. ex. -timeout=10")

	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println(err)
	}

	r := csv.NewReader(file)

	done := make(chan bool)

	go qa(r, done)

	select {
	case <-done:
	case <-time.After(time.Duration(*timeout) * time.Second):
		fmt.Println("\nout of time")
	}

	fmt.Println("Thank you for playing!")
}

func qa(r *csv.Reader, done chan bool) {
	var questionCounter, correctCounter int
	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}

		fmt.Printf("Q: %s = ", rec[0])
		answer := rec[1]

		reader := bufio.NewReader(os.Stdin)
		userAnswer, _ := reader.ReadString('\n')

		if strings.ToLower(strings.Trim(userAnswer, "\n")) == strings.ToLower(strings.Trim(answer, "\n")) {
			fmt.Println("Correct")
			correctCounter++
		} else {
			fmt.Println("Wrong")
		}

		questionCounter++
	}

	fmt.Printf("You have answered %d out of %d questions correctly!", correctCounter, questionCounter)

	done <- true
}
