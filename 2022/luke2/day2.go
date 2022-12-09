package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Round struct {
	opponent string
	me       string
	score    int
}

func main() {
	f, _ := os.Open("luke2/input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	totalScore, totalScoreDay2 := 0, 0
	for scanner.Scan() {
		round := strings.Split(scanner.Text(), " ")
		r := Round{rockPaperScissors(round[0]), rockPaperScissors(round[1]), 0}                                     //This looks
		r2 := Round{rockPaperScissors(round[0]), howToPlay(rockPaperScissors(round[0]), methodToPlay(round[1])), 0} //Awful
		totalScore += calculateScore(r)
		totalScoreDay2 += calculateScore(r2)
	}
	fmt.Println("Part 1", totalScore)
	fmt.Println("Part 2", totalScoreDay2)
}

func calculateScore(r Round) int {
	r.score += getValOfChoice(r.me)
	r.score += playRound(r.opponent, r.me)
	return r.score
}

func playRound(o string, m string) int {
	if o == "Rock" && m == "Scissor" {
		return 0
	} else if o == "Paper" && m == "Rock" {
		return 0
	} else if o == "Scissor" && m == "Paper" {
		return 0
	} else if m == o {
		return 3
	}
	return 6
}

func getValOfChoice(c string) int {
	switch c {
	case "Rock":
		return 1
	case "Paper":
		return 2
	case "Scissor":
		return 3
	default:
		return 0
	}
}

func rockPaperScissors(s string) string {
	switch s {
	case "A", "X":
		return "Rock"
	case "B", "Y":
		return "Paper"
	case "C", "Z":
		return "Scissor"
	default:
		return ""
	}
}

func methodToPlay(s string) string {
	switch s {
	case "X":
		return "Lose"
	case "Y":
		return "Tie"
	case "Z":
		return "Win"
	default:
		return ""
	}
}

func howToPlay(inp string, method string) string {
	switch {
	case inp == "Rock" && method == "Win", inp == "Scissor" && method == "Lose", inp == "Paper" && method == "Tie":
		return "Paper"
	case inp == "Rock" && method == "Tie", inp == "Scissor" && method == "Win", inp == "Paper" && method == "Lose":
		return "Rock"
	case inp == "Rock" && method == "Lose", inp == "Scissor" && method == "Tie", inp == "Paper" && method == "Win":
		return "Scissor"
	default:
		return ""
	}
}
