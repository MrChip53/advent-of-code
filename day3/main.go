package main

import (
	"AdventOfCode2023/pkg/input"
	"fmt"
	"strconv"
)

type PartNumber struct {
	Started   bool
	StartX    int
	EndX      int
	NumberStr string
}

func (p *PartNumber) AddNumber(x int, ch int32) {
	if p.NumberStr == "" {
		p.StartX = x
		p.Started = true
	}
	p.NumberStr += string(ch)
}

func (p *PartNumber) HasStarted() bool {
	return p.Started
}

func (p *PartNumber) GetNumber() int {
	n, err := strconv.Atoi(p.NumberStr)
	if err != nil {
		panic(err)
	}
	return n
}

func (p *PartNumber) IsPartNumber(lines []string) bool {
	if len(lines) != 3 {
		panic("need 3 lines")
	}

	numLen := len(p.NumberStr)

	firstChar := p.StartX
	if firstChar > 0 {
		firstChar--
	}

	lastChar := p.StartX + numLen - 1
	if lastChar < len(lines[1])-1 {
		lastChar++
	}

	for i := firstChar; i <= lastChar; i++ {
		ch1 := lines[0][i]
		ch2 := lines[1][i]
		ch3 := lines[2][i]

		if isSymbol(ch1) || isSymbol(ch2) || isSymbol(ch3) {
			return true
		}
	}

	return false
}

func isSymbol(ch uint8) bool {
	return ch != 46 && (ch-'0' < 0 || ch-'0' > 9)
}

func isNumber(ch int) bool {
	return ch-'0' >= 0 && ch-'0' <= 9
}

func getBlankLine(len int) string {
	var line string
	for i := 0; i < len; i++ {
		line += "."
	}
	return line
}

func part1(lines []string) int {
	var partNumSum int

	for y, line := range lines {
		currentNumber := PartNumber{}
		for x, ch := range line {
			chNum := ch - '0'
			if chNum >= 0 && chNum <= 9 {
				currentNumber.AddNumber(x, ch)
			}
			if x == len(line)-1 || chNum < 0 || chNum > 9 {
				if currentNumber.HasStarted() {
					var pLines []string
					if y != 0 {
						pLines = append(pLines, lines[y-1])
					} else {
						pLines = append(pLines, getBlankLine(len(line)))
					}
					pLines = append(pLines, line)
					if y != len(lines)-1 {
						pLines = append(pLines, lines[y+1])
					} else {
						pLines = append(pLines, getBlankLine(len(line)))
					}
					if currentNumber.IsPartNumber(pLines) {
						partNumSum += currentNumber.GetNumber()
					}
					currentNumber = PartNumber{}
				}
			}
		}
	}

	return partNumSum
}

func getNumberCount(str [3]uint8) ([]int, int) {
	var numCount int
	var offsets []int
	if isNumber(int(str[0])) {
		numCount++
		offsets = append(offsets, -1)
		if !isNumber(int(str[1])) && isNumber(int(str[2])) {
			numCount++
			offsets = append(offsets, 1)
		}
	} else if isNumber(int(str[1])) {
		numCount++
		offsets = append(offsets, 0)
	} else if isNumber(int(str[2])) {
		numCount++
		offsets = append(offsets, 1)
	}
	return offsets, numCount
}

func getSurroundingNumbersCount(x int, lines []string) ([][]int, int) {
	var nums int

	var numLocations [][]int
	for _, line := range lines {
		var sArr [3]uint8

		startPos := x - 1
		endPos := x + 2

		if x == 0 {
			startPos = x
		} else if x == len(line)-1 {
			endPos = x
		}

		copy(sArr[:], line[startPos:endPos])

		offsets, count := getNumberCount(sArr)
		nums += count

		numLocations = append(numLocations, offsets)
	}

	return numLocations, nums
}

func getNumber(x int, str string) int {
	if !isNumber(int(str[x])) {
		return 0
	}

	var start = x
	for i := x; i >= 0; i-- {
		if isNumber(int(str[i])) {
			start = i
		} else {
			break
		}
	}

	currentNumber := PartNumber{}
	for i := start; i < len(str); i++ {
		if !isNumber(int(str[i])) {
			break
		}
		currentNumber.AddNumber(i, int32(str[i]))
	}
	return currentNumber.GetNumber()
}

func part2(lines []string) int {
	var gearSum int

	for y, line := range lines {
		for x, ch := range line {
			if ch == '*' {
				var pLines []string
				if y != 0 {
					pLines = append(pLines, lines[y-1])
				} else {
					pLines = append(pLines, getBlankLine(len(line)))
				}
				pLines = append(pLines, line)
				if y != len(lines)-1 {
					pLines = append(pLines, lines[y+1])
				} else {
					pLines = append(pLines, getBlankLine(len(line)))
				}
				loc, count := getSurroundingNumbersCount(x, pLines)
				if count == 2 {
					var gearNumbers []int
					for idx, l := range loc {
						for _, k := range l {
							gearNumbers = append(gearNumbers, getNumber(x+k, pLines[idx]))
						}
					}
					gearSum += gearNumbers[0] * gearNumbers[1]
				}
			}
		}
	}

	return gearSum
}

func main() {
	file, scanner := input.OpenInputText("./input/day3.txt")
	defer file.Close()

	var lines []string
	for scanner.Scan() {
		str := scanner.Text()
		lines = append(lines, str)
	}

	p1 := part1(lines)
	p2 := part2(lines)

	fmt.Printf("Part 1 Answer: %d\n", p1)
	fmt.Printf("Part 2 Answer: %d\n", p2)
}
