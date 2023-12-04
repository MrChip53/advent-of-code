package main

import (
	"AdventOfCode2023/pkg/input"
	"fmt"
)

var textNumbers = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func part1() int {
	file, scanner := input.OpenInputText("./input/day1.txt")
	defer file.Close()

	a := 0

	for scanner.Scan() {
		f := -1
		l := -1

		str := scanner.Text()
		strLen := len(str)

		for i := 0; i < strLen; i++ {
			if f == -1 {
				if str[i]-'0' > 0 && str[i]-'0' <= 9 {
					f = int(str[i] - '0')
				}
			}
			if l == -1 {
				if str[strLen-1-i]-'0' > 0 && str[strLen-1-i]-'0' <= 9 {
					l = int(str[strLen-1-i] - '0')
				}
			}
			if l != -1 && f != -1 {
				break
			}
		}
		a += f*10 + l
	}

	return a
}

func part2() int {
	var letterMap = make(map[uint8]bool)

	for _, str := range textNumbers {
		letterMap[str[0]] = true
	}

	file, scanner := input.OpenInputText("./input/day1.txt")
	defer file.Close()

	a := 0

	for scanner.Scan() {
		f := -1
		l := -1

		str := scanner.Text()
		strLen := len(str)

		for i := 0; i < strLen; i++ {
			if f == -1 {
				if str[i]-'0' > 0 && str[i]-'0' <= 9 {
					f = int(str[i] - '0')
				} else if letterMap[str[i]] {
					fI := checkNumString(str, i)
					if fI != -1 {
						f = fI
					}
				}
			}
			if l == -1 {
				if str[strLen-1-i]-'0' > 0 && str[strLen-1-i]-'0' <= 9 {
					l = int(str[strLen-1-i] - '0')
				} else if letterMap[str[strLen-1-i]] {
					lI := checkNumString(str, strLen-1-i)
					if lI != -1 {
						l = lI
					}
				}
			}
			if l != -1 && f != -1 {
				break
			}
		}
		a += f*10 + l
	}

	return a
}

func checkNumString(str string, pos int) int {
	for idx, numStr := range textNumbers {

		isThisItMan := true
		for i := 0; i < len(numStr); i++ {
			if pos+i >= len(str) {
				isThisItMan = false
				break
			}
			if numStr[i] != str[pos+i] {
				isThisItMan = false
				break
			}
		}
		if isThisItMan {
			return idx + 1
		}
	}

	return -1
}

func main() {
	p1 := part1()
	p2 := part2()

	fmt.Printf("Part 1 Answer: %d\n", p1)
	fmt.Printf("Part 2 Answer: %d\n", p2)
}
