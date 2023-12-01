package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Config struct {
	day int
	part int
	data Input
	extraParams []any
}

func NewConfig(day int, part int, data Input, extraParams ...any) *Config {
	return &Config{day, part, data, extraParams}
}

func ParseConfig() *Config {
	var day, part int
	var inputData string
	flag.IntVar(&day, "d", 1, "the day of the challenge")
	flag.IntVar(&day, "day", 1, "the day of the challenge")
	flag.IntVar(&part, "p", 1, "the part of the challenge")
	flag.IntVar(&part, "part", 1, "the part of the challenge")
	flag.StringVar(&inputData, "i", "real", "the input data to use")
	flag.StringVar(&inputData, "input", "real", "the input data to use")
	flag.Parse()
	fmt.Println(os.Args)
	var data Input
	if day < 1 || day > 25 {
		fmt.Println("Day must be between 1 and 25")
	}
	if part < 1 || part > 2 {
		fmt.Println("Part must be 1 or 2")
		return nil
	}
	if inputData == "real" {
		data = *NewRealInput()
	} else {
		testNumber, err := strconv.Atoi(inputData)
		if err != nil {
			fmt.Println("Input must be either 'real' or a number")
			return nil
		} else {
			data = *NewTestInput(testNumber)
		}
	}
	args := flag.Args()
	extraParams := make([]any, len(args))
	for i, arg := range args {
		extraParams[i] = arg
	}
	return NewConfig(day, part, data, extraParams...)
}

func (c *Config) GetDay() int {
	return c.day
}

func (c *Config) GetExtraParams() []any {
	return c.extraParams
}

func (c *Config) GetInputType() string {
	if c.data.real {
		return "real"
	} else {
		return "test" + strconv.Itoa(c.data.test)
	}
}

func (c *Config) GetInputData() string {
	fileName := c.getInputFilename()
	fileName, _ = filepath.Abs(fileName)
	fmt.Println(fileName)
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}
	return string(data)
}

func (c *Config) GetPart() int {
	return c.part
}

func (c *Config) getInputFilename() string {
	return fmt.Sprintf("inputs/day%d/data%s.txt", c.day, c.data.String())
}

type ConfigForTest struct {
	Config
}

func NewConfigForTest(c *Config) *ConfigForTest {
	return &ConfigForTest{*c}
}

func (c *ConfigForTest) GetInputData() string {
	fileName := c.getInputFilename()
	fileName, _ = filepath.Abs(fileName)
	fmt.Println(fileName)
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}
	return string(data)
}

func (c *ConfigForTest) getInputFilename() string {
	return fmt.Sprintf("../../inputs/day%d/data%s.txt", c.day, c.data.String())
}