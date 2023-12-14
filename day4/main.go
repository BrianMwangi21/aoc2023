package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strings"
	"time"
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

func sanitizeData(data []string) []string {
	var result []string

	for _, v := range data {
		index := strings.Index(v, ":")
		result = append(result, v[index+2:])
	}

	return result
}

func extractDigits(text string) []string {
	pattern := `\b\d+`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(text, -1)
	return matches
}

func getMatchedNumbers(combo []string) []string {
	var result []string
	winning_numbers := extractDigits(combo[0])
	sample_numbers := extractDigits(combo[1])

	for _, w := range winning_numbers {
		for _, s := range sample_numbers {
			if w == s {
				result = append(result, w)
			}
		}
	}

	return result
}

func partOne(data []string) int {
	var result int
	sanitized := sanitizeData(data)

	for _, v := range sanitized {
		combo := strings.Split(v, " | ")
		matched_numbers := getMatchedNumbers(combo)

		if len(matched_numbers) > 0 {
			if len(matched_numbers) == 1 {
				result += 1
			} else {
				result += int(math.Pow(float64(2), float64(len(matched_numbers)-1)))
			}
		}
	}

	return result
}

func partTwo(data []string) int {
	now := time.Now()
	var copies []string
	sanitized := sanitizeData(data)

	copies = append(copies, sanitized...)
	idx := 0

	for idx != len(copies) {
		combo := strings.Split(copies[idx], " | ")
		matched_numbers := getMatchedNumbers(combo)

		if len(matched_numbers) > 0 {
			idx_match := 0
			for idx_inner := range sanitized {
				if copies[idx] == sanitized[idx_inner] {
					idx_match = idx_inner
					break
				}
			}

			for count := range matched_numbers {
				copies = append(copies, sanitized[idx_match+1+count])
			}
		}

		idx++
	}

	fmt.Println("Time taken: ", time.Since(now))
	return len(copies)
}

func partTwoWithMap(data []string) int {
	now := time.Now()
	match_map := make(map[string]int)
	var copies []string
	sanitized := sanitizeData(data)

	for index, v := range sanitized {
		match_map[v] = index
	}

	copies = append(copies, sanitized...)
	idx := 0

	for idx != len(copies) {
		combo := strings.Split(copies[idx], " | ")
		matched_numbers := getMatchedNumbers(combo)

		if len(matched_numbers) > 0 {
			idx_match := match_map[copies[idx]]
			for count := range matched_numbers {
				copies = append(copies, sanitized[idx_match+1+count])
			}
		}

		idx++
	}

	fmt.Println("Time taken: ", time.Since(now))
	return len(copies)
}

func main() {
	data := readFile("input.txt")

	fmt.Println("Part One: ", partOne(data))
	fmt.Println("Part Two: ", partTwo(data))
	fmt.Println("Part Two With Map: ", partTwoWithMap(data))
}
