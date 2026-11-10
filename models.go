package main

import "time"

// Question stores information about a quiz question.
type Question struct {
	Question   string   `json:"question"`
	Options    []string `json:"options"`
	Answer     int      `json:"answer"`
	Difficulty string   `json:"difficulty"`
}

// Attempt stores the result of a completed quiz.
type Attempt struct {
	Score          int       `json:"score"`
	TotalQuestions int       `json:"total_questions"`
	Difficulty     string    `json:"difficulty"`
	Date           time.Time `json:"date"`
}
