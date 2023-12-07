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

func guard(e error) {
	if e != nil {
		panic(e)
	}
}

func sanitizeData(data []string) []string {
	var result []string

	for _, v := range data {
		index := strings.Index(v, ":")
		result = append(result, v[index+2:])
	}

	return result
}

func extractTotalValues(text string) (int, int, int) {
	blue, red, green := 0, 0, 0

	regexPattern := `\b\d+ (red|blue|green)\b`
	re := regexp.MustCompile(regexPattern)
	matches := re.FindAllString(text, -1)
	fmt.Println(matches)

	for _, m := range matches {
		split := strings.Split(m, " ")

		fmt.Println(split[0])
		fmt.Println(split[1])

		num, err := strconv.Atoi(split[0])
		guard(err)

		switch split[1] {
		case "red":
			red += num
		case "green":
			green += num
		case "blue":
			blue += num
		default:
			fmt.Println("Unknown color")
		}
	}

	return red, green, blue
}

func partOne(data []string) int {
	var possible []int
	redLimit, greenLimit, blueLimit := 12, 13, 14

	sanitized := sanitizeData(data)

	for index, v := range sanitized {
		red, green, blue := extractTotalValues(v)

		if red <= redLimit && green <= greenLimit && blue <= blueLimit {
			fmt.Println(index+1, ":", red, "red", green, "green", blue, "blue", "- possible")
			possible = append(possible, index+1)
		}

		break
	}

	fmt.Println("Possible:", possible)

	return 0
}

func main() {
	data := readFile("input.txt")

	fmt.Println(partOne(data))
}
