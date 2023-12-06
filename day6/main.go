package main

import (
	"AdventOfCode2023/pkg/input"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseLinesPt1(lines []string) ([]int, []int) {
	var times []int
	var distances []int

	for _, line := range lines {
		if strings.HasPrefix(line, "Time:") {
			timeStr := strings.TrimSpace(strings.ReplaceAll(line, "Time:", ""))
			timesSplit := strings.Split(timeStr, " ")
			for _, t := range timesSplit {
				if t == "" {
					continue
				}
				ti, err := strconv.Atoi(t)
				if err != nil {
					panic(err)
				}
				times = append(times, ti)
			}
		} else if strings.HasPrefix(line, "Distance:") {
			dStr := strings.TrimSpace(strings.ReplaceAll(line, "Distance:", ""))
			dSplit := strings.Split(dStr, " ")
			for _, d := range dSplit {
				if d == "" {
					continue
				}
				di, err := strconv.Atoi(d)
				if err != nil {
					panic(err)
				}
				distances = append(distances, di)
			}
		}
	}
	return times, distances
}

func parseLinesPt2(lines []string) (int, int) {
	var timeI int
	var distance int

	for _, line := range lines {
		if strings.HasPrefix(line, "Time:") {
			timeStr := strings.ReplaceAll(strings.ReplaceAll(line, "Time:", ""), " ", "")
			ti, err := strconv.Atoi(timeStr)
			if err != nil {
				panic(err)
			}
			timeI = ti
		} else if strings.HasPrefix(line, "Distance:") {
			dStr := strings.ReplaceAll(strings.ReplaceAll(line, "Distance:", ""), " ", "")
			di, err := strconv.Atoi(dStr)
			if err != nil {
				panic(err)
			}
			distance = di
		}
	}
	return timeI, distance
}

func part1(lines []string) int {
	times, distances := parseLinesPt1(lines)

	var acc []int
	for idx := range times {
		var numWays int
		for i := 0; i < times[idx]; i++ {
			distance := (times[idx] - i) * i
			if distance > distances[idx] {
				numWays++
			}
		}
		if numWays > 0 {
			acc = append(acc, numWays)
		}
	}

	var num = 1
	for _, a := range acc {
		num *= a
	}
	return num
}

func part2(lines []string) int {
	times, distances := parseLinesPt2(lines)

	var numWays int
	for i := 0; i < times; i++ {
		distance := (times - i) * i
		if distance > distances {
			numWays++
		}
	}
	return numWays
}

func main() {
	file, scanner := input.OpenInputText("./input/day6.txt")
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
