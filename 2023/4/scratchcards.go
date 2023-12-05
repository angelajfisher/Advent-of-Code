package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"regexp"
	"strings"
	"math"
)

type scratchcard struct {
	id			int
	winningNums	[]int
	cardNums	[]int
}

func main() {
	lines, _ := readLines()
	scratchcards := parseData(lines)

	points := 0
	for _, scratchcard := range scratchcards {
		wins := countWins(scratchcard)
		points += int(math.Pow(2, wins-1))
	}

	fmt.Printf("The answer to part one is: %v\n", points)
}

func countWins(scratchcard scratchcard) float64 {
	var wins float64 = 0
	for _, num := range scratchcard.winningNums {
		for _, n := range scratchcard.cardNums {
			if n == num {
				wins++
				continue
			}
		}
	}
	return wins
}

// parseData takes an array of strings of data and
// turns them into an array of scratchcard structs
func parseData(lines []string) []scratchcard {
	scratchcards := []scratchcard{}
	for i := 0; i < len(lines); i++ {
		var scratchcard scratchcard

		regex := regexp.MustCompile(`\d+`)
		idIndex := regex.FindStringIndex(lines[i])

		scratchcard.id, _ = strconv.Atoi(lines[i][idIndex[0]:idIndex[1]])

		line := lines[i][idIndex[1]+2:]

		allNums := strings.Split(line, "|")
		winngingNums := regex.FindAllString(allNums[0], -1)
		cardNums := regex.FindAllString(allNums[1], -1)

		for i := 0; i < len(winngingNums); i++ {
			num, _ := strconv.Atoi(winngingNums[i])
			scratchcard.winningNums = append(scratchcard.winningNums, num)
		}
		for i := 0; i < len(cardNums); i++ {
			num, _ := strconv.Atoi(cardNums[i])
			scratchcard.cardNums = append(scratchcard.cardNums, num)
		}
		scratchcards = append(scratchcards, scratchcard)
	}
	return scratchcards
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines() ([]string, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
