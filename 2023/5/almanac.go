package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type almanac struct {
	seeds                 []*seed
	seedToSoilMap         [][]int
	soilToFertilizerMap   [][]int
	fertilizerToWaterMap  [][]int
	waterToLightMap       [][]int
	lightToTempMap        [][]int
	tempToHumidityMap     [][]int
	humidityToLocationMap [][]int
}

type seed struct {
	id         int
	soil       int
	fertilizer int
	water      int
	light      int
	temp       int
	humidity   int
	location   int
}

func main() {
	lines, _ := readLines()

	almanac := mapSeeds(dataToAlmanac(lines))

	location := almanac.seeds[0].location
	for _, seedRef := range almanac.seeds {
		seed := *seedRef
		if location > seed.location {
			location = seed.location
		}
	}
	fmt.Printf("The answer to part one is: %v\n", location)
}

// mapSeeds takes in an almanac and populates the seeds within it
// with their mapped values according to the almanac, then returns
// the updated almanac
func mapSeeds(almanac almanac) almanac {
	for i, seedRef := range almanac.seeds {
		seed := *seedRef

		seed.soil = convertFromMap(seed.id, almanac.seedToSoilMap)
		seed.fertilizer = convertFromMap(seed.soil, almanac.soilToFertilizerMap)
		seed.water = convertFromMap(seed.fertilizer, almanac.fertilizerToWaterMap)
		seed.light = convertFromMap(seed.water, almanac.waterToLightMap)
		seed.temp = convertFromMap(seed.light, almanac.lightToTempMap)
		seed.humidity = convertFromMap(seed.temp, almanac.tempToHumidityMap)
		seed.location = convertFromMap(seed.humidity, almanac.humidityToLocationMap)

		almanac.seeds[i] = &seed
	}
	return almanac
}

func convertFromMap(convertNum int, ranges [][]int) int {
	convertedNum := convertNum

	for _, mapping := range ranges {
		if convertNum >= mapping[1] && convertNum <= mapping[1]+mapping[2] {
			convertedNum = (mapping[0] + mapping[2]) - (mapping[1] + mapping[2] - convertNum)
			break
		}
	}
	return convertedNum
}

// dataToAlmanac takes a slice of lines (strings)
// and parses them into an almanac struct
func dataToAlmanac(lines []string) almanac {
	almanac := almanac{}
	index := 0

	seedLines, i := parseSection(lines)
	index += i
	almanac.seedToSoilMap, i = parseSection(lines[index:])
	index += i
	almanac.soilToFertilizerMap, i = parseSection(lines[index:])
	index += i
	almanac.fertilizerToWaterMap, i = parseSection(lines[index:])
	index += i
	almanac.waterToLightMap, i = parseSection(lines[index:])
	index += i
	almanac.lightToTempMap, i = parseSection(lines[index:])
	index += i
	almanac.tempToHumidityMap, i = parseSection(lines[index:])
	index += i
	almanac.humidityToLocationMap, _ = parseSection(lines[index:])

	for _, num := range seedLines[0] {
		newSeed := seed{id: num}
		almanac.seeds = append(almanac.seeds, &newSeed)
	}
	return almanac
}

// parseSection takes a slice of lines and returns a slice of
// slices of numbers corresponding to the contents of each line,
// as well as the index of the end of the section
func parseSection(lines []string) (section [][]int, end int) {
	section = [][]int{}
	regex := regexp.MustCompile(`\d+`)

	for i, line := range lines {
		nums := []int{}
		strNums := regex.FindAllString(line, -1)

		if len(strNums) == 0 {
			end = i + 2 // offset by 2 to account for lines w/o rows
			break
		}

		for _, n := range strNums {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}

		section = append(section, nums)
	}
	return
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
