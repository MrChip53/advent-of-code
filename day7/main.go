package main

import (
	"AdventOfCode2023/pkg/input"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	HAND_HIGH_CARD = iota
	HAND_ONE_PAIR
	HAND_TWO_PAIR
	HAND_THREE_KIND
	HAND_FULL_HOUSE
	HAND_FOUR_KIND
	HAND_FIVE_KIND
)

var cardspt1 = map[string]int{
	"A": 14,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 11,
	"Q": 12,
	"K": 13,
}

var cardspt2 = map[string]int{
	"A": 14,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"J": 1,
	"Q": 12,
	"K": 13,
}

type Hand struct {
	Cards    string
	HandType int
	Bid      int
}

func readHands(lines []string, pt2 bool) []*Hand {
	var hands []*Hand

	for _, line := range lines {
		var handType = HAND_HIGH_CARD
		split := strings.Split(strings.TrimSpace(line), " ")
		bidNum, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}

		splitCards := strings.Split(split[0], "")
		var cardMap = make(map[string]int)
		for _, c := range splitCards {
			if _, ok := cardMap[c]; !ok {
				cardMap[c] = 0
			}

			cardMap[c]++
		}

		if pt2 && cardMap["J"] > 0 {
			var highestCard string
			var count int
			for idx, c := range cardMap {
				if idx == "J" {
					continue
				}
				if c > count {
					highestCard = idx
					count = c
				}
			}
			cardMap[highestCard] += cardMap["J"]
		}

		for idx, c := range cardMap {
			if idx == "J" && pt2 {
				continue
			}
			if c == 2 {
				if handType == HAND_HIGH_CARD {
					handType = HAND_ONE_PAIR
				} else if handType == HAND_ONE_PAIR {
					handType = HAND_TWO_PAIR
				} else if handType == HAND_THREE_KIND {
					handType = HAND_FULL_HOUSE
				}
			} else if c == 3 {
				if handType == HAND_ONE_PAIR {
					handType = HAND_FULL_HOUSE
				} else {
					handType = HAND_THREE_KIND
				}
			} else if c == 4 {
				handType = HAND_FOUR_KIND
			} else if c == 5 {
				handType = HAND_FIVE_KIND
			}
		}

		hands = append(hands, &Hand{
			Cards:    split[0],
			HandType: handType,
			Bid:      bidNum,
		})
	}

	return hands
}

func part1(lines []string) int {
	hands := readHands(lines, false)

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].HandType != hands[j].HandType {
			return hands[i].HandType < hands[j].HandType
		}

		for k := range hands[i].Cards {
			if hands[i].Cards[k] != hands[j].Cards[k] {
				return cardspt1[string(hands[i].Cards[k])] < cardspt1[string(hands[j].Cards[k])]
			}
		}

		return false
	})

	var ret = 0
	for i, hand := range hands {
		ret += hand.Bid * (i + 1)
	}

	return ret
}

func part2(lines []string) int {
	hands := readHands(lines, true)

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].HandType != hands[j].HandType {
			return hands[i].HandType < hands[j].HandType
		}

		for k := range hands[i].Cards {
			if hands[i].Cards[k] != hands[j].Cards[k] {
				return cardspt2[string(hands[i].Cards[k])] < cardspt2[string(hands[j].Cards[k])]
			}
		}

		return false
	})

	var ret = 0
	for i, hand := range hands {
		ret += hand.Bid * (i + 1)
	}

	return ret
}

func main() {
	file, scanner := input.OpenInputText("./input/day7.txt")
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
