package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// RunQuiz starts and manages a quiz.
func RunQuiz(questions []Question, difficulty string, reader *bufio.Reader) (Attempt, error) {

	// Select questions based on difficulty.
	var selectedQuestions []Question

	for _, q := range questions {
		if strings.EqualFold(q.Difficulty, difficulty) {
			selectedQuestions = append(selectedQuestions, q)
		}
	}

	// Check if questions are available.
	if len(selectedQuestions) == 0 {
		fmt.Println("No questions available for this difficulty.")
		return Attempt{}, nil
	}

	score := 0

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("          QUIZ STARTED")
	fmt.Println("================================")
	fmt.Println("Difficulty:", strings.ToUpper(difficulty))
	fmt.Println("Total Questions:", len(selectedQuestions))

	// Display questions one by one.
	for i, q := range selectedQuestions {

		fmt.Println()
		fmt.Printf("Question %d: %s\n", i+1, q.Question)

		for j, option := range q.Options {
			fmt.Printf("%c. %s\n", 'A'+j, option)
		}

		// Accept and validate the answer.
		var answer int

		for {
			fmt.Print("Your answer (A/B/C/D): ")

			input, err := reader.ReadString('\n')

			if err != nil && len(input) == 0 {
				return Attempt{}, fmt.Errorf("could not read answer: %w", err)
			}

			input = strings.ToUpper(strings.TrimSpace(input))

			if len(input) == 1 && input[0] >= 'A' && input[0] <= 'D' {
				answer = int(input[0]-'A') + 1
			} else {
				number, err := strconv.Atoi(input)

				if err != nil || number < 1 || number > len(q.Options) {
					fmt.Println("Invalid input! Please enter A, B, C or D.")
					continue
				}

				answer = number
			}

			break
		}

		// Check the answer.
		if answer == q.Answer {
			fmt.Println("Correct answer!")
			score++
		} else {
			fmt.Println("Incorrect answer!")
			fmt.Printf("Correct answer: %c. %s\n",
				'A'+q.Answer-1, q.Options[q.Answer-1])
		}
	}

	// Display final results.
	fmt.Println()
	fmt.Println("================================")
	fmt.Println("           QUIZ RESULT")
	fmt.Println("================================")

	fmt.Printf("Score: %d/%d\n", score, len(selectedQuestions))

	percentage := float64(score) / float64(len(selectedQuestions)) * 100

	fmt.Printf("Percentage: %.2f%%\n", percentage)

	if percentage >= 80 {
		fmt.Println("Performance: Excellent!")
	} else if percentage >= 50 {
		fmt.Println("Performance: Good job!")
	} else {
		fmt.Println("Performance: Keep practising!")
	}

	fmt.Println("================================")

	// Create an attempt record.
	attempt := Attempt{
		Score:          score,
		TotalQuestions: len(selectedQuestions),
		Difficulty:     difficulty,
		Date:           time.Now(),
	}

	return attempt, nil
}
