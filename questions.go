package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// LoadQuestions reads questions from a JSON file.
func LoadQuestions(filename string) ([]Question, error) {

	// Open and read the file.
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, fmt.Errorf("could not read question file: %w", err)
	}

	// Check if the file is empty.
	if len(strings.TrimSpace(string(data))) == 0 {
		return nil, fmt.Errorf("question file is empty")
	}

	// Convert JSON data into Question structures.
	var questions []Question

	err = json.Unmarshal(data, &questions)

	if err != nil {
		return nil, fmt.Errorf("invalid JSON format: %w", err)
	}

	// Check if the question list is empty.
	if len(questions) == 0 {
		return nil, fmt.Errorf("no questions found in the file")
	}

	// Validate each question.
	for i, q := range questions {

		if strings.TrimSpace(q.Question) == "" {
			return nil, fmt.Errorf("question %d has no question text", i+1)
		}

		if len(q.Options) != 4 {
			return nil, fmt.Errorf("question %d must have exactly 4 options", i+1)
		}

		if q.Answer < 1 || q.Answer > len(q.Options) {
			return nil, fmt.Errorf("question %d has an invalid answer number", i+1)
		}

		difficulty := strings.ToLower(strings.TrimSpace(q.Difficulty))

		if difficulty != "easy" &&
			difficulty != "medium" &&
			difficulty != "hard" {
			return nil, fmt.Errorf("question %d has an invalid difficulty", i+1)
		}
	}

	return questions, nil
}
