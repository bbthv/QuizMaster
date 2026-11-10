package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	// Load questions from JSON.
	questions, err := LoadQuestions("questions.json")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Questions loaded successfully!")
	fmt.Println("Total questions:", len(questions))

	// Main menu.
	for {

		fmt.Println()
		fmt.Println("================================")
		fmt.Println("        WELCOME TO QUIZMASTER")
		fmt.Println("================================")
		fmt.Println("1. Start Quiz")
		fmt.Println("2. View Quiz History")
		fmt.Println("3. View Leaderboard")
		fmt.Println("4. Exit")
		fmt.Println("================================")

		fmt.Print("Enter your choice: ")

		choice, err := reader.ReadString('\n')

		if err != nil && len(choice) == 0 {
			fmt.Println("\nInput ended. Exiting QuizMaster.")
			return
		}

		choice = strings.TrimSpace(choice)

		switch choice {

		case "1":

			var difficulty string

			for {

				fmt.Println()
				fmt.Println("Select Difficulty:")
				fmt.Println("1. Easy")
				fmt.Println("2. Medium")
				fmt.Println("3. Hard")

				fmt.Print("Enter your choice: ")

				input, err := reader.ReadString('\n')

				if err != nil && len(input) == 0 {
					fmt.Println("\nInput ended. Exiting QuizMaster.")
					return
				}

				input = strings.TrimSpace(input)

				switch input {

				case "1":
					difficulty = "easy"

				case "2":
					difficulty = "medium"

				case "3":
					difficulty = "hard"

				default:
					fmt.Println("Invalid choice! Try again.")
					continue
				}

				break
			}

			// Start the quiz and receive its result.
			attempt, err := RunQuiz(questions, difficulty, reader)

			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			// Save the completed attempt.
			if attempt.TotalQuestions > 0 {

				err = SaveAttempt(attempt)

				if err != nil {
					fmt.Println("Could not save quiz history:", err)
				} else {
					fmt.Println("Quiz result saved successfully!")
				}
			}

		case "2":

			// Display previous quiz attempts.
			DisplayHistory()

		case "3":

			fmt.Println("Leaderboard module is coming soon!")

		case "4":

			fmt.Println("Thank you for using QuizMaster!")
			return

		default:

			fmt.Println("Invalid choice! Please select 1-4.")
		}
	}
}
