package day1

import (
	"aoc_2023/utils/arrays"
	"aoc_2023/utils/maps"
	"aoc_2023/utils/stringutils"
	"strconv"
	"strings"
)

type Solver struct{}

func (*Solver) SolvePart1(lines []string, extraParams ...any) string {
	cr := NewBasicCalibrationReader(lines)
	sum := cr.getCalibrationSum()
	return strconv.Itoa(sum)
}

func (*Solver) SolvePart2(lines []string, extraParams ...any) string {
	cr := NewExtraCalibrationReader(lines)
	sum := cr.getCalibrationSum()
	return strconv.Itoa(sum)
}

var extraDigits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

type CalibrationReader struct {
	mapping map[string]int
	lines   []string
}

func NewBasicCalibrationReader(lines []string) *CalibrationReader {
	return &CalibrationReader{
		mapping: make(map[string]int),
		lines:   lines,
	}
}

func NewExtraCalibrationReader(lines []string) *CalibrationReader {
	return &CalibrationReader{
		mapping: extraDigits,
		lines:   lines,
	}
}

func (cr *CalibrationReader) getCalibrationSum() int {
	calibrations := arrays.Map(cr.lines, cr.getCalibration)
	return arrays.Sum(calibrations)
}

func (cr *CalibrationReader) getCalibration(line string) int {
	return cr.calculateCalibrationValue(cr.findFirstDigit(line), cr.findLastDigit(line))
}

func (cr *CalibrationReader) findFirstDigit(line string) int {
	for i := 0; i < len(line); i++ {
		if stringutils.IsDigit(rune(line[i])) {
			return int(line[i] - '0')
		}
		value, ok := cr.getExtraDigitPrefix(line[i:])
		if ok {
			return value
		}
	}
	return -1
}

func (cr *CalibrationReader) findLastDigit(line string) int {
	for i := 0; i < len(line); i++ {
		if stringutils.IsDigit(rune(line[len(line)-i-1])) {
			return int(line[len(line)-i-1] - '0')
		}
		value, ok := cr.hasExtraDigitSuffix(line[:len(line)-i])
		if ok {
			return value
		}
	}
	return -1
}

func (cr *CalibrationReader) getExtraDigitPrefix(line string) (int, bool) {
	return maps.FindValue(cr.mapping, func(key string, _ int) bool {
		return strings.HasPrefix(line, key)
	})
}

func (cr *CalibrationReader) hasExtraDigitSuffix(line string) (int, bool) {
	return maps.FindValue(cr.mapping, func(key string, _ int) bool {
		return strings.HasSuffix(line, key)
	})
}

func (cr *CalibrationReader) calculateCalibrationValue(firstDigit, lastDigit int) int {
	return firstDigit*10 + lastDigit
}
