package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	id   int
	sets []set
}

type set struct {
	blue  int
	red   int
	green int
}

func main() {
	lines, _ := readLines()

	games := []game{}
	maxSet := set{blue: 14, red: 12, green: 13}

	for i := 0; i < len(lines); i++ {
		games = append(games, parseData(lines[i]))
	}

	possibleGames := validateGames(games, maxSet)
	sumPossible := 0
	for i := 0; i < len(possibleGames); i++ {
		sumPossible += possibleGames[i].id
	}

	minimumCubes := maxCubes(games)
	totalPower := 0
	for i := 0; i < len(minimumCubes); i++ {
		totalPower += (minimumCubes[i].blue * minimumCubes[i].red * minimumCubes[i].green)
	}

	fmt.Printf("The answer to part one is: %v\n", sumPossible)
	fmt.Printf("The answer to part two is: %v\n", totalPower)
}

// parseData takes a string of data
// and turns it into a game struct
func parseData(line string) game {
	var game game

	regex := regexp.MustCompile(`\d+`)
	idIndex := regex.FindStringIndex(line)

	game.id, _ = strconv.Atoi(line[idIndex[0]:idIndex[1]])

	line = line[idIndex[1]+2:]

	sets := strings.Split(line, ";")
	for i := 0; i < len(sets); i++ {
		colors := strings.Split(sets[i], ",")
		var set = set{}

		for j := 0; j < len(colors); j++ {
			currentColor := colors[j]
			if strings.Contains(currentColor, "blue") {
				set.blue, _ = strconv.Atoi(regex.FindString(currentColor))
				continue
			}
			if strings.Contains(currentColor, "red") {
				set.red, _ = strconv.Atoi(regex.FindString(currentColor))
				continue
			}
			if strings.Contains(currentColor, "green") {
				set.green, _ = strconv.Atoi(regex.FindString(currentColor))
				continue
			}
		}
		game.sets = append(game.sets, set)
	}
	return game
}

// maxCubes takes a list of games and returns a list
// of sets corresponding to the minimum number of
// cubes required to play the given game
func maxCubes(games []game) []set {
	maximums := []set{}

	for i := 0; i < len(games); i++ {
		maxCubes := set{red: 0, blue: 0, green: 0}
		for j := 0; j < len(games[i].sets); j++ {
			if games[i].sets[j].blue > maxCubes.blue {
				maxCubes.blue = games[i].sets[j].blue
			}
			if games[i].sets[j].red > maxCubes.red {
				maxCubes.red = games[i].sets[j].red
			}
			if games[i].sets[j].green > maxCubes.green {
				maxCubes.green = games[i].sets[j].green
			}
		}
		maximums = append(maximums, maxCubes)
	}
	return maximums
}


// validateGames compares the sets in each game to the max possible
// set and returns those that fit within the constraints
func validateGames(games []game, maxSet set) []game {
	possibleGames := []game{}

	for i := 0; i < len(games); i++ {
		currentGame := games[i]
		possible := true
		for j := 0; j < len(currentGame.sets); j++ {
			if currentGame.sets[j].blue > maxSet.blue || currentGame.sets[j].red > maxSet.red || currentGame.sets[j].green > maxSet.green {
				possible = false
			}
		}
		if possible {
			possibleGames = append(possibleGames, currentGame)
		}
	}
	return possibleGames
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
