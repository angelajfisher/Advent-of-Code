package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines, _ := readLines()

	for i := 0; i < len(lines); i++ {
		fmt.Println(lines[i])
	}

	fmt.Printf("Part one solution: %v\n", digitCalibrationValues(lines))

	fmt.Printf("Part two solution: %v\n", allCalibrationValues(lines))
}

// allCalibrationValues produces the solution to part two, where
// values can be both words and digits, and returns the sum
func allCalibrationValues(lines []string) int {
	sum := 0

	for i := 0; i < len(lines); i++ {
		nums := getAllNums(lines[i])

		firstNum := nums[0]
		lastNum := nums[len(nums)-1]

		if len(lastNum) > 1 {
			lastNum = confirmLastNum(lines[i])
		}

		numStr := wordToNum(firstNum) + wordToNum(lastNum)
		num, _ := strconv.Atoi(numStr)

		sum += num
	}
	return sum
}

// confirmLastNum returns the written-out number at the end
// of the given string
func confirmLastNum(line string) string {
	slice := ""
	for i := len(line) - 1; i >= 0; i-- {

		slice = string(line[i]) + slice

		num := wordToNum(slice)
		if num != slice {
			return num
		}
	}
	return ""
}

// strToDigit converts a stringified number to an integer,
// whether it's a word or an arabic numeral
func wordToNum(strNum string) string {
	if strings.HasPrefix(strNum, "one") {
		return "1"
	}
	if strings.HasPrefix(strNum, "two") {
		return "2"
	}
	if strings.HasPrefix(strNum, "three") {
		return "3"
	}
	if strings.HasPrefix(strNum, "four") {
		return "4"
	}
	if strings.HasPrefix(strNum, "five") {
		return "5"
	}
	if strings.HasPrefix(strNum, "six") {
		return "6"
	}
	if strings.HasPrefix(strNum, "seven") {
		return "7"
	}
	if strings.HasPrefix(strNum, "eight") {
		return "8"
	}
	if strings.HasPrefix(strNum, "nine") {
		return "9"
	}

	return strNum
}

// digitCalibrationValues produces the solution to part one
// and returns the sum of all the calibration values as digits
func digitCalibrationValues(lines []string) int {
	sum := 0

	for i := 0; i < len(lines); i++ {
		regex := regexp.MustCompile(
			`\d{1}`,
		)
		nums := regex.FindAllString(lines[i], -1)

		numStr := nums[0] + nums[len(nums)-1]
		num, _ := strconv.Atoi(numStr)

		sum += num
	}
	return sum
}

// getAllNums compares the given string to regex that extracts
// all of the numbers (written-out or arabic) and returns them
// all in an array
func getAllNums(line string) []string {
	regex := regexp.MustCompile(
		`\d{1}|(one|two|three|four|five|six|seven|eight|nine)`,
	)

	return regex.FindAllString(line, -1)
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
