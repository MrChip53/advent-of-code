package main

import (
	"AdventOfCode2023/pkg/input"
	"fmt"
	"strconv"
	"strings"
)

type GamePull struct {
	RedCubes   int
	BlueCubes  int
	GreenCubes int
}

type GameEntry struct {
	Id    int
	Pulls []GamePull
}

func (g *GameEntry) isValid(red int, green int, blue int) bool {
	for _, p := range g.Pulls {
		if p.RedCubes > red || p.GreenCubes > green || p.BlueCubes > blue {
			return false
		}
	}
	return true
}

func (g *GameEntry) getCubePower() int {
	red := 0
	green := 0
	blue := 0
	for _, p := range g.Pulls {
		if p.RedCubes > red {
			red = p.RedCubes
		}
		if p.BlueCubes > blue {
			blue = p.BlueCubes
		}
		if p.GreenCubes > green {
			green = p.GreenCubes
		}
	}

	return red * green * blue
}

func parseGameEntry(str string) GameEntry {
	// Example str input
	// Game 1: 1 blue; 4 green, 5 blue; 11 red, 3 blue, 11 green; 1 red, 10 green, 4 blue; 17 red, 12 green, 7 blue; 3 blue, 19 green, 15 red
	firstSplit := strings.Split(str, ":")
	gamePart := firstSplit[0]
	pullPart := firstSplit[1]

	gamePart = strings.ReplaceAll(gamePart, "Game ", "")
	gamePart = strings.TrimSpace(gamePart)

	gameId, err := strconv.Atoi(gamePart)
	if err != nil {
		panic(err)
	}

	pulls := strings.Split(pullPart, ";")

	var gamePulls []GamePull
	for _, pull := range pulls {
		gamePull := GamePull{}
		pull = strings.TrimSpace(pull)
		splitPull := strings.Split(pull, ",")
		for _, p := range splitPull {
			p = strings.TrimSpace(p)
			split := strings.Split(p, " ")
			pullNum, err := strconv.Atoi(split[0])
			if err != nil {
				panic(err)
			}

			switch split[1] {
			case "red":
				gamePull.RedCubes = pullNum
			case "green":
				gamePull.GreenCubes = pullNum
			case "blue":
				gamePull.BlueCubes = pullNum
			}
		}
		gamePulls = append(gamePulls, gamePull)
	}

	return GameEntry{
		Id:    gameId,
		Pulls: gamePulls,
	}
}

func part1() int {
	file, scanner := input.OpenInputText("./input/day2.txt")
	defer file.Close()

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	var games []GameEntry
	for scanner.Scan() {
		str := scanner.Text()

		entry := parseGameEntry(str)
		games = append(games, entry)
	}

	sumIds := 0
	for _, game := range games {
		if game.isValid(maxRed, maxGreen, maxBlue) {
			sumIds += game.Id
		}
	}

	return sumIds
}

func part2() int {
	file, scanner := input.OpenInputText("./input/day2.txt")
	defer file.Close()

	var games []GameEntry
	for scanner.Scan() {
		str := scanner.Text()

		entry := parseGameEntry(str)
		games = append(games, entry)
	}

	sumPower := 0
	for _, game := range games {
		sumPower += game.getCubePower()
	}

	return sumPower
}

func main() {
	p1 := part1()
	p2 := part2()

	fmt.Printf("Part 1 Answer: %d\n", p1)
	fmt.Printf("Part 2 Answer: %d\n", p2)
}
