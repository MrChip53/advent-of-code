package main

import (
	"AdventOfCode2023/pkg/input"
	"fmt"
	"strings"
	"time"
)

const (
	LEFT = iota
	RIGHT
)

type Map struct {
	CurKeys      []string
	Count        int
	Instructions []int
	Forks        map[string][]string
}

func primeFactors(n int) map[int]int {
	factors := make(map[int]int)

	for n%2 == 0 {
		factors[2]++
		n /= 2
	}

	for i := 3; i*i <= n; i += 2 {
		for n%i == 0 {
			factors[i]++
			n /= i
		}
	}

	if n > 2 {
		factors[n]++
	}

	return factors
}

func getLCM(nums []int) int {
	var maps []map[int]int
	for _, v := range nums {
		maps = append(maps, primeFactors(v))
	}

	factors := make(map[int]int)
	for _, v := range maps {
		for f, p := range v {
			factors[f] = max(factors[f], p)
		}
	}

	lcm := 1
	for f, p := range factors {
		for i := 0; i < p; i++ {
			lcm *= f
		}
	}

	return lcm
}

func readMap(lines []string) *Map {
	var myMap = &Map{
		Forks: make(map[string][]string),
	}
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		if i == 0 {
			for _, l := range line {
				if l == 'R' {
					myMap.Instructions = append(myMap.Instructions, RIGHT)
				} else {
					myMap.Instructions = append(myMap.Instructions, LEFT)
				}
			}
			continue
		}

		if strings.Contains(line, "=") {
			line = strings.ReplaceAll(line, "(", "")
			line = strings.ReplaceAll(line, ")", "")
			splits := strings.Split(line, " = ")
			key := splits[0]
			forks := strings.Split(splits[1], ", ")
			myMap.Forks[key] = forks
			if strings.HasSuffix(key, "A") {
				myMap.CurKeys = append(myMap.CurKeys, key)
			}
		}
	}
	return myMap
}

func part1(lines []string) int {
	var steps int

	myMap := readMap(lines)
	numI := len(myMap.Instructions)

	var curKey = "AAA"
	for curKey != "ZZZ" {
		sideIdx := (steps%numI + numI) % numI
		side := myMap.Instructions[sideIdx]
		curKey = myMap.Forks[curKey][side]
		steps++
	}
	return steps
}

func part2(lines []string) int {
	var steps int

	myMap := readMap(lines)
	numI := len(myMap.Instructions)

	var stepArr []int

	for len(stepArr) != len(myMap.CurKeys) {
		sideIdx := (steps%numI + numI) % numI
		side := myMap.Instructions[sideIdx]
		for i, k := range myMap.CurKeys {
			myMap.CurKeys[i] = myMap.Forks[k][side]
			if strings.HasSuffix(myMap.CurKeys[i], "Z") {
				stepArr = append(stepArr, steps+1)
			}
		}
		steps++
	}
	return getLCM(stepArr)
}

func main() {
	file, scanner := input.OpenInputText("./input/day8.txt")
	defer file.Close()

	var lines []string
	for scanner.Scan() {
		str := scanner.Text()
		lines = append(lines, str)
	}

	var start time.Time
	var end time.Time
	start = time.Now()
	p1 := part1(lines)
	end = time.Now()
	fmt.Printf("Part 1 completed in %v\n", end.Sub(start))

	start = time.Now()
	p2 := part2(lines)
	end = time.Now()
	fmt.Printf("Part 2 completed in %v\n", end.Sub(start))

	fmt.Printf("Part 1 Answer: %d\n", p1)
	fmt.Printf("Part 2 Answer: %d\n", p2)
}
