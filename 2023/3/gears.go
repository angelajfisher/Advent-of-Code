package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type schematic struct {
	symbols []int
	gearSymbols []int
	numberIndexes [][]int
	numbers []string
}

func main() {
	lines, _ := readLines()
	lastIndex := len(lines[0])-1

	schematics := parseLines(lines)

	partNums := grabParts(schematics, lastIndex)
	sum := 0
	for i := 0; i < len(partNums); i++ {
		sum += partNums[i]
	}
	fmt.Printf("The answer to part one is: %v\n", sum)

	ratio := 0
	for _, num := range grabGears(schematics, lastIndex) {
		ratio += num
	}
	fmt.Printf("The answer to part two is: %v", ratio)
}

// This spaghetti takes all the schematics and returns an array
// of ratios for the actual gear parts
func grabGears(schematics []schematic, lastIndex int) []int {
	gears := []int{}

	for i := 0; i < len(schematics); i++ {
		for _, index := range schematics[i].gearSymbols {
			nearby := []int{}

			nums := neighboringNums(index, lastIndex, schematics[i].numberIndexes)
			if nums[0] >= 0 {
				num, _ := strconv.Atoi(schematics[i].numbers[nums[0]])
				nearby = append(nearby, num)
				if nums[1] >= 0 {
					num, _ := strconv.Atoi(schematics[i].numbers[nums[1]])
					nearby = append(nearby, num)
				}
			}
			if i != 0 {
				nums = neighboringNums(index, lastIndex, schematics[i-1].numberIndexes)
				if nums[0] >= 0 {
					num, _ := strconv.Atoi(schematics[i-1].numbers[nums[0]])
					nearby = append(nearby, num)
					if nums[1] >= 0 {
						num, _ := strconv.Atoi(schematics[i-1].numbers[nums[1]])
						nearby = append(nearby, num)
					}
				}
			}
			if i != 0 {
				nums = neighboringNums(index, lastIndex, schematics[i+1].numberIndexes)
				if nums[0] >= 0 {
					num, _ := strconv.Atoi(schematics[i+1].numbers[nums[0]])
					nearby = append(nearby, num)
					if nums[1] >= 0 {
						num, _ := strconv.Atoi(schematics[i+1].numbers[nums[1]])
						nearby = append(nearby, num)
					}
				}
			}
			if len(nearby) > 1 {
				gears = append(gears, (nearby[0] * nearby[1]))
			}
		}
	}
	return gears
}

// neighboringNums takes a line's start and end values for each num
// (as an array of arrays) as well as the index of a gear symbol and
// the final index of the line, search area, then returns up to two
// indexes of neighboring numbers as an array
func neighboringNums(index int, lastIndex int, numberIndexes [][]int) [2]int {
	start := index - 1
	end := index + 1
	gearNumIdx := [2]int{-1, -1}

	for i := 0; i < len(numberIndexes); i++ {
		current := numberIndexes[i]

		if (current[0] >= start && current[0] <= end) || (current[1]-1 >= start && current[1]-1 <= end) {
			if gearNumIdx[0] > -1 {
				gearNumIdx[1] = i
				return gearNumIdx
			} else {
				gearNumIdx[0] = i
			}
		}
	}
	return gearNumIdx
}

// grabParts takes all of the schematics and returns the
// numbers within them that correspond to parts.
func grabParts(schematics []schematic, lastIndex int) []int {
	parts := []int{}

	for i := 0; i < len(schematics); i++ {
		current := schematics[i].numberIndexes
		for j := 0; j < len(current); j++ {
			var start int
			var end int

			if current[j][0] > 0 {
				start = current[j][0]-1
			} else {
				start = 0
			}
			if current[j][1] <= lastIndex {
				end = current[j][1]
			} else {
				end = lastIndex
			}

			checkSchematics := [][]int{schematics[i].symbols,}
			if i != 0 {
				checkSchematics = append(checkSchematics, schematics[i-1].symbols)
			}
			if i != len(schematics)-1 {
				checkSchematics = append(checkSchematics, schematics[i+1].symbols)
			}

			if checkForParts(start, end, checkSchematics) {
				num, _ := strconv.Atoi(schematics[i].numbers[j])
				parts = append(parts, num)
			}
		}
	}
	return parts
}

// checkForParts takes multiple lines' symbols (as an array of arrays)
// as well as the start and end indexes corresponding to a number's
// search area, then returns true if the number is adjacent to a
// symbol and returns false if it is not.
func checkForParts(start int, end int, symbols [][]int) bool {
	for i := 0; i < len(symbols); i++ {
		for j := start; j <= end; j++ {
			for _, index := range symbols[i] {
				if index == j {
					return true
				}
			}
		}
	}
	return false
}

// parseLines takes an array of lines and parses them
// into schematic structs that track the indexes of
// numbers and symbols in a given line.
func parseLines(lines []string) []schematic {
	schematics := []schematic{}

	for i := 0; i < len(lines); i++ {
		line := schematic{}

		numRegex := regexp.MustCompile(`\d+`)
		line.numberIndexes = numRegex.FindAllStringIndex(lines[i], -1)
		line.numbers = numRegex.FindAllString(lines[i], -1)

		symbRegex := regexp.MustCompile(`[^\d^.]`)
		symbols := symbRegex.FindAllStringIndex(lines[i], -1)
		for j := 0; j < len(symbols); j++ {
			line.symbols = append(line.symbols, symbols[j][0])
		}

		gearRegex := regexp.MustCompile((`\*+`))
		gears := gearRegex.FindAllStringIndex(lines[i], -1)
		for j := 0; j < len(gears); j++ {
			line.gearSymbols = append(line.gearSymbols, gears[j][0])
		}

		schematics = append(schematics, line)
	}
	return schematics
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
