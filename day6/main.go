package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Record struct {
	time     int
	distance int
}

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
		log.Fatal(err)
	}
}

func getMatches(text string, part string) []int {
	var digits []int
	var digitString string

	pattern := `\b\d+`
	r := regexp.MustCompile(pattern)
	matches := r.FindAllString(text, -1)

	if part == "1" {
		for _, v := range matches {
			digit, err := strconv.Atoi(v)
			guard(err)

			digits = append(digits, digit)
		}
	} else if part == "2" {
		for _, v := range matches {
			digitString += v
		}

		digit, err := strconv.Atoi(digitString)
		guard(err)

		digits = append(digits, digit)
	}

	return digits
}

func getRecords(data []string, part string) []Record {
	var records []Record
	var times, distances []int

	times = getMatches(data[0], part)
	distances = getMatches(data[1], part)

	for i := 0; i < len(times); i++ {
		records = append(records, Record{times[i], distances[i]})
	}

	return records
}

func getWaysToWin(record Record) int {
	var waysToWin []int
	time := record.time
	distance := record.distance

	for i := 1; i < time; i++ {
		timeRemaining := time - i
		distanceCovered := timeRemaining * i

		if distanceCovered > distance {
			waysToWin = append(waysToWin, i)
		}
	}

	return len(waysToWin)
}

func partOneAndTwo(data []string, part string) int {
	result := 1
	records := getRecords(data, part)

	for _, v := range records {
		waysToWin := getWaysToWin(v)
		result *= waysToWin
	}

	return result
}

func main() {
	data := readFile("input.txt")

	fmt.Println("Part One:", partOneAndTwo(data, "1"))
	fmt.Println("Part Two:", partOneAndTwo(data, "2"))
}
