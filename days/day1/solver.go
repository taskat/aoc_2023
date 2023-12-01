package day1

import (
	"strconv"
	"strings"
)

type Solver struct{}

func mapDigits(r rune) rune {
	if isDigit(r) {
		return r
	}
	return -1
}

func sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func (*Solver) SolvePart1(input string, extraParams ...any) string {
	lines := strings.Split(input, "\n")
	numbers := make([]int, len(lines))
	for i, line := range lines {
		digits := strings.Map(mapDigits, line)
		calibrationString := string(append([]byte{digits[0]}, digits[len(digits)-1]))
		calibration, err := strconv.Atoi(calibrationString)
		if err != nil {
			panic(err)
		}
		numbers[i] = calibration
	}
	sum := sum(numbers)
	return strconv.Itoa(sum)
}

func (*Solver) SolvePart2(input string, extraParams ...any) string {
	lines := strings.Split(input, "\n")
	numbers := make([]int, len(lines))
	for i, line := range lines {
		line = addExtraDigits(line)
		digits := strings.Map(mapDigits, line)
		calibrationString := string(append([]byte{digits[0]}, digits[len(digits)-1]))
		calibration, err := strconv.Atoi(calibrationString)
		if err != nil {
			panic(err)
		}
		numbers[i] = calibration
	}
	sum := sum(numbers)
	return strconv.Itoa(sum)
}

func addExtraDigits(line string) string {
	extraDigits := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	firstDigit, lasDigist := false, false
	for i := 0; i < len(line)-3; i++ {
		if isDigit(rune(line[i])) {
			firstDigit = true
		}
		if isDigit(rune(line[len(line)-i-1])) {
			lasDigist = true
		}
		for key, value := range extraDigits {
			if strings.HasPrefix(line[i:], key) && !firstDigit {
				line = strings.Replace(line, key, value, 1)
			}
			if strings.HasSuffix(line[:len(line)-i], key) && !lasDigist {
				keyLength := len(key)
				line = line[:len(line)-i-keyLength] + value + line[len(line)-i:]
			}
		}
	}
	return line
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
