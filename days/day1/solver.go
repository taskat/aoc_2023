package day1

import (
	"aoc_2023/utils/arrays"
	"aoc_2023/utils/stringutils"
	"strconv"
	"strings"
)

type Solver struct{}

func mapDigits(r rune) rune {
	if stringutils.IsDigit(r) {
		return r
	}
	return -1
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
	sum := arrays.Sum(numbers)
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
	sum := arrays.Sum(numbers)
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
		if stringutils.IsDigit(rune(line[i])) {
			firstDigit = true
		}
		if stringutils.IsDigit(rune(line[len(line)-i-1])) {
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
