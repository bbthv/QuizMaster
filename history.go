package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadHistory reads previous quiz attempts.
func LoadHistory() ([]Attempt, error) {

	data, err := os.ReadFile("history.json")

	// If history file does not exist, return an empty history.
	if os.IsNotExist(err) {
		return []Attempt{}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("could not read history: %w", err)
	}

	// Handle an empty history file.
	if len(data) == 0 {
		return []Attempt{}, nil
	}

	var attempts []Attempt

	err = json.Unmarshal(data, &attempts)

	if err != nil {
		return nil, fmt.Errorf("invalid history file: %w", err)
	}

	return attempts, nil
}

// SaveAttempt saves a completed quiz result.
func SaveAttempt(attempt Attempt) error {

	// Load previous attempts.
	attempts, err := LoadHistory()

	if err != nil {
		return err
	}

	// Add the new attempt.
	attempts = append(attempts, attempt)

	// Convert attempts into JSON.
	data, err := json.MarshalIndent(attempts, "", "    ")

	if err != nil {
		return fmt.Errorf("could not convert history to JSON: %w", err)
	}

	// Save the updated history.
	err = os.WriteFile("history.json", data, 0644)

	if err != nil {
		return fmt.Errorf("could not save history: %w", err)
	}

	return nil
}

// DisplayHistory displays all previous quiz attempts.
func DisplayHistory() {

	attempts, err := LoadHistory()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("          QUIZ HISTORY")
	fmt.Println("================================")

	if len(attempts) == 0 {
		fmt.Println("No previous quiz attempts found.")
		return
	}

	for i, attempt := range attempts {

		fmt.Printf("\nAttempt %d\n", i+1)

		fmt.Println("Score:",
			attempt.Score, "/", attempt.TotalQuestions)

		fmt.Println("Difficulty:",
			attempt.Difficulty)

		fmt.Println("Date:",
			attempt.Date.Format("02 Jan 2006, 03:04 PM"))
	}

	fmt.Println()
	fmt.Println("================================")
}
