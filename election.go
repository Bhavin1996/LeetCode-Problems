package main

import "fmt"

func findElectionWinner(votes [][]string) string {
	score := make(map[string]int)

	// Calculate scores based on voter preferences
	for _, vote := range votes {
		for i, candidate := range vote {
			score[candidate] += 3 - i
		}
	}

	maxScore := 0
	winner := ""

	// Find the candidate with the highest score
	for candidate, candidateScore := range score {
		if candidateScore > maxScore {
			maxScore = candidateScore
			winner = candidate
		}
	}

	return winner
}

func main() {
	votes := [][]string{
		{"Raj", "Amit", "Kriti"},
		{"Raj", "Vinod", "Sam"},
		{"Amit", "Raj", "Kriti"},
	}

	winner := findElectionWinner(votes)
	fmt.Printf("The winner is: %s\n", winner)
}
