package main

import (
	"fmt"
	"sort"
)

// DisplayLeaderboard shows the highest scoring quiz attempts.
func DisplayLeaderboard() {

	// Load previous quiz attempts.
	attempts, err := LoadHistory()

	if err != nil {
		fmt.Println("Error loading leaderboard:", err)
		return
	}

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("           LEADERBOARD")
	fmt.Println("================================")

	if len(attempts) == 0 {
		fmt.Println("No quiz attempts available yet.")
		return
	}

	// Sort attempts by percentage.
	sort.Slice(attempts, func(i, j int) bool {

		percentage1 := float64(attempts[i].Score) /
			float64(attempts[i].TotalQuestions)

		percentage2 := float64(attempts[j].Score) /
			float64(attempts[j].TotalQuestions)

		return percentage1 > percentage2
	})

	// Display only the top 5 attempts.
	limit := len(attempts)

	if limit > 5 {
		limit = 5
	}

	for i := 0; i < limit; i++ {

		attempt := attempts[i]

		percentage := float64(attempt.Score) /
			float64(attempt.TotalQuestions) * 100

		fmt.Printf("\nRank: %d\n", i+1)

		fmt.Printf("Score: %d/%d\n",
			attempt.Score, attempt.TotalQuestions)

		fmt.Printf("Percentage: %.2f%%\n", percentage)

		fmt.Println("Difficulty:", attempt.Difficulty)

		fmt.Println("Date:",
			attempt.Date.Format("02 Jan 2006, 03:04 PM"))

		fmt.Println("--------------------------------")
	}

	fmt.Println("================================")
}
