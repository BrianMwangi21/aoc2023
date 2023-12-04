package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func getDigits(value string) int {
	first, last := 0, 0

	for i := 0; i < len(value); i++ {
		number, err := strconv.Atoi(string(value[i]))

		if err != nil {
			// Do nothing
		} else {
			if first == 0 {
				first = number
			}
			last = number
		}
	}

	digit, err := strconv.Atoi(fmt.Sprintf("%d%d", first, last))

	if err != nil {
		// Do nothing
	}

	return digit
}

func getVariations(value string) string {
	digit_map := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}
	first_index, first_number, last_index, last_number := 100, 100, 0, 0

	for i := 1; i < 10; i++ {
		// Replace values
		num_s := fmt.Sprintf("%d", i)
		new_value := strings.ReplaceAll(value, digit_map[i], num_s)
		index := strings.Index(new_value, num_s)

		if index != -1 && index < first_index {
			first_index = index
			first_number = i
		}

		if index != -1 && index > last_index {
			last_index = index
			last_number = i
		}
	}

	final_value := strings.ReplaceAll(value, digit_map[first_number], fmt.Sprintf("%d", first_number))
	return strings.ReplaceAll(final_value, digit_map[last_number], fmt.Sprintf("%d", last_number))
}

func partOne(data []string) int {
	var result int

	for v := range data {
		digit := getDigits(data[v])
		result += digit
	}

	return result
}

func partTwo(data []string) int {
	var result int

	for v := range data {
		value := getVariations(data[v])
		digit := getDigits(value)
		result += digit
	}

	return result
}

func main() {
	data := readFile("input.txt")

	fmt.Println("Part One:", partOne(data))
	fmt.Println("Part Two:", partTwo(data))
}
