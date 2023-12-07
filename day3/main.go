package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile(filename string) []string {
	var data []string

	f, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func guard(err error) {
	if err != nil {
		panic(err)
	}
}

func getNumbers(text string) []string {
	regexPattern := `\b\d+`
	re := regexp.MustCompile(regexPattern)
	matches := re.FindAllString(text, -1)
	return matches
}

func build_region(data []string, row int, num_idx int, digit string) []string {
	var (
		upper string
		mid   string
		lower string
	)

	end := num_idx + len(digit)

	// Upper
	if row == 0 {
		upper = strings.Repeat(".", len(digit)+2)
	} else {
		if num_idx == 0 {
			upper = "." + data[row-1][num_idx:end]
		} else {
			upper = data[row-1][num_idx-1 : end]
		}

		if end < len(data[row]) {
			upper += data[row-1][end : end+1]
		} else {
			upper += "."
		}
	}

	// Mid
	if num_idx == 0 {
		mid = "." + digit
	} else {
		mid = data[row][num_idx-1:num_idx] + digit
	}

	if end < len(data[row]) {
		mid += data[row][end : end+1]
	} else {
		mid += "."
	}

	// Lower
	if row == len(data)-1 {
		lower = strings.Repeat(".", len(digit)+2)
	} else {
		if num_idx == 0 {
			lower = "." + data[row+1][num_idx:end]
		} else {
			lower = data[row+1][num_idx-1 : end]
		}

		if end < len(data[row]) {
			lower += data[row+1][end : end+1]
		} else {
			lower += "."
		}
	}

	region := []string{upper, mid, lower}
	return region
}

func checkIfPartNumber(region []string) bool {
	var result bool
	regexPattern := `[^0-9.]`
	re := regexp.MustCompile(regexPattern)

	for _, text := range region {
		matches := re.FindAllString(text, -1)
		if len(matches) > 0 {
			result = true
		}
	}

	return result
}

func partOne(data []string) int {
	var result int

	for row, text := range data {
		digits := getNumbers(text)

		if len(digits) > 0 {
			for _, digit := range digits {
				num_idx := strings.Index(text, digit)
				region := build_region(data, row, num_idx, digit)
				isPartNumber := checkIfPartNumber(region)

				if isPartNumber {
					num, err := strconv.Atoi(digit)
					guard(err)
					result += num
				}
			}
		}
	}

	return result
}

func main() {
	data := readFile("input.test.txt")

	fmt.Println("Part One:", partOne(data))
}
