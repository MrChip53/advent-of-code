package main

import (
	"AdventOfCode2023/pkg/input"
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	Id             int
	WinningNumbers map[int]bool
	Numbers        []int
	Matches        int
}

func (c *Card) GetValue() int {
	var value int
	for _, num := range c.Numbers {
		if c.WinningNumbers[num] {
			if value > 0 {
				value = value * 2
			} else {
				value = 1
			}
		}
	}
	return value
}

func NewCard(str string) *Card {
	str = strings.ReplaceAll(str, "  ", " ")

	split1 := strings.Split(str, ":")
	cardId, err := strconv.Atoi(strings.TrimSpace(strings.ReplaceAll(split1[0], "Card", "")))
	if err != nil {
		panic(err)
	}
	split2 := strings.Split(split1[1], "|")
	winningNumSplit := strings.Split(strings.TrimSpace(split2[0]), " ")

	var wMap = make(map[int]bool)
	for _, v := range winningNumSplit {
		v = strings.TrimSpace(v)
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		wMap[i] = true
	}

	myNumbers := strings.Split(strings.TrimSpace(split2[1]), " ")

	var nums []int
	for _, v := range myNumbers {
		v = strings.TrimSpace(v)
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		nums = append(nums, i)
	}

	var matches int
	for _, v := range nums {
		if wMap[v] {
			matches++
		}
	}

	return &Card{
		Id:             cardId,
		WinningNumbers: wMap,
		Numbers:        nums,
		Matches:        matches,
	}
}

func part1(lines []string) int {
	var value int
	for _, l := range lines {
		card := NewCard(l)
		value += card.GetValue()
	}
	return value
}

type CardCounter struct {
	index int
	queue []int
	Value int
}

func (c *CardCounter) Add(id int) {
	c.queue = append(c.queue, id)
}

func (c *CardCounter) Next() bool {
	if c.index >= len(c.queue) {
		return false
	}

	c.Value = c.queue[c.index]
	c.index++
	return true
}

func part2(lines []string) int {
	var cards = make(map[int]*Card)
	var numCards int

	var cardCounter CardCounter

	for _, l := range lines {
		card := NewCard(l)
		cards[card.Id] = card
		cardCounter.Add(card.Id)
	}

	var numCardsMap = make(map[int]int)
	for cardCounter.Next() {
		cId := cardCounter.Value
		numCardsMap[cId]++
		numCards++
		card := cards[cId]
		if card.Matches > 0 {
			for i := 1; i <= card.Matches; i++ {
				if i <= len(lines) {
					cardCounter.Add(cId + i)
				}
			}
		}
	}

	return numCards
}

func main() {
	file, scanner := input.OpenInputText("./input/day4.txt")
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
