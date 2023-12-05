package main

import (
	"AdventOfCode2023/pkg/input"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

const (
	MAP_MODE_SOIL = iota
	MAP_MODE_FERT
	MAP_MODE_WATER
	MAP_MODE_LIGHT
	MAP_MODE_TEMP
	MAP_MODE_HUMIDITY
	MAP_MODE_LOCATION
)

type SeedValues struct {
	Seed     int
	Soil     int
	Fert     int
	Water    int
	Light    int
	Temp     int
	Humidity int
	Location int
}

type MapValues struct {
	Source      int
	Destination int
	Range       int
}

type SeedRange struct {
	Start int
	Range int
}

type SeedMaps struct {
	Seeds    []int
	SeedsPt2 []SeedRange

	SeedSoil   []*MapValues
	SoilFert   []*MapValues
	FertWater  []*MapValues
	WaterLight []*MapValues
	LightTemp  []*MapValues
	TempHumi   []*MapValues
	HumiLoc    []*MapValues

	Mode int
}

func (s *SeedMaps) getDestValue(source int, values []*MapValues) int {
	for _, v := range values {
		if source >= v.Source && source <= v.Source+v.Range-1 {
			return v.Destination + source - v.Source
		}
	}
	return -1
}

func (s *SeedMaps) getSourceValue(dest int, values []*MapValues) int {
	for _, v := range values {
		if dest >= v.Destination && dest <= v.Destination+v.Range-1 {
			return v.Source + dest - v.Destination
		}
	}
	return -1
}

func (s *SeedMaps) getSeedFromLoc(loc int) int {
	humi := s.getSourceValue(loc, s.HumiLoc)
	if humi == -1 {
		humi = loc
	}
	temp := s.getSourceValue(humi, s.TempHumi)
	if temp == -1 {
		temp = humi
	}
	light := s.getSourceValue(temp, s.LightTemp)
	if light == -1 {
		light = temp
	}
	water := s.getSourceValue(light, s.WaterLight)
	if water == -1 {
		water = light
	}
	fert := s.getSourceValue(water, s.FertWater)
	if fert == -1 {
		fert = water
	}
	soil := s.getSourceValue(fert, s.SoilFert)
	if soil == -1 {
		soil = fert
	}
	seed := s.getSourceValue(soil, s.SeedSoil)
	if seed == -1 {
		seed = soil
	}

	for _, sm := range s.SeedsPt2 {
		if seed >= sm.Start && seed <= sm.Start+sm.Range-1 {
			return seed
		}
	}

	return -1
}

func (s *SeedMaps) getSeedValues(seed int) *SeedValues {
	soil := s.getDestValue(seed, s.SeedSoil)
	if soil == -1 {
		soil = seed
	}
	fert := s.getDestValue(soil, s.SoilFert)
	if fert == -1 {
		fert = soil
	}
	water := s.getDestValue(fert, s.FertWater)
	if water == -1 {
		water = fert
	}
	light := s.getDestValue(water, s.WaterLight)
	if light == -1 {
		light = water
	}
	temp := s.getDestValue(light, s.LightTemp)
	if temp == -1 {
		temp = light
	}
	humi := s.getDestValue(temp, s.TempHumi)
	if humi == -1 {
		humi = temp
	}
	loc := s.getDestValue(humi, s.HumiLoc)
	if loc == -1 {
		loc = humi
	}

	return &SeedValues{
		Seed:     seed,
		Soil:     soil,
		Fert:     fert,
		Water:    water,
		Light:    light,
		Temp:     temp,
		Humidity: humi,
		Location: loc,
	}
}

func (s *SeedMaps) parseSeedsPt1(line string) {
	line = strings.TrimSpace(strings.ReplaceAll(line, "seeds:", ""))
	numStrs := strings.Split(line, " ")

	for i := range numStrs {
		num, err := strconv.Atoi(numStrs[i])
		if err != nil {
			panic(err)
		}
		s.Seeds = append(s.Seeds, num)
	}
}

func (s *SeedMaps) parseSeedsPt2(line string) {
	line = strings.TrimSpace(strings.ReplaceAll(line, "seeds:", ""))
	numStrs := strings.Split(line, " ")

	if len(numStrs)%2 != 0 {
		panic("odd number")
	}

	for i := 0; i < len(numStrs); i += 2 {
		num1, err := strconv.Atoi(numStrs[i])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(numStrs[i+1])
		if err != nil {
			panic(err)
		}
		s.SeedsPt2 = append(s.SeedsPt2, SeedRange{Start: num1, Range: num2})
	}
}

func (s *SeedMaps) addToMap(key int, value int, l int) {
	switch s.Mode {
	case MAP_MODE_SOIL:
		s.SeedSoil = append(s.SeedSoil, &MapValues{
			Source:      key,
			Destination: value,
			Range:       l,
		})
	case MAP_MODE_FERT:
		s.SoilFert = append(s.SoilFert, &MapValues{
			Source:      key,
			Destination: value,
			Range:       l,
		})
	case MAP_MODE_WATER:
		s.FertWater = append(s.FertWater, &MapValues{
			Source:      key,
			Destination: value,
			Range:       l,
		})
	case MAP_MODE_LIGHT:
		s.WaterLight = append(s.WaterLight, &MapValues{
			Source:      key,
			Destination: value,
			Range:       l,
		})
	case MAP_MODE_TEMP:
		s.LightTemp = append(s.LightTemp, &MapValues{
			Source:      key,
			Destination: value,
			Range:       l,
		})
	case MAP_MODE_HUMIDITY:
		s.TempHumi = append(s.TempHumi, &MapValues{
			Source:      key,
			Destination: value,
			Range:       l,
		})
	case MAP_MODE_LOCATION:
		s.HumiLoc = append(s.HumiLoc, &MapValues{
			Source:      key,
			Destination: value,
			Range:       l,
		})
	}
}

func (s *SeedMaps) parseMapLine(line string) {
	numStrs := strings.Split(line, " ")
	if len(numStrs) != 3 {
		panic("not 3")
	}

	valNumStart, err := strconv.Atoi(numStrs[0]) // The source
	if err != nil {
		panic(err)
	}
	keyNumStart, err := strconv.Atoi(numStrs[1])
	if err != nil {
		panic(err)
	}
	l, err := strconv.Atoi(numStrs[2])
	if err != nil {
		panic(err)
	}

	s.addToMap(keyNumStart, valNumStart, l)
}

func NewSeedMap(lines []string, pt1 bool) *SeedMaps {
	var seedMap = &SeedMaps{}

	for _, line := range lines {
		if strings.HasPrefix(line, "seeds:") {
			if pt1 {
				seedMap.parseSeedsPt1(line)
			} else {
				seedMap.parseSeedsPt2(line)
			}
		} else if strings.HasPrefix(line, "seed-to-soil") {
			seedMap.Mode = MAP_MODE_SOIL
		} else if strings.HasPrefix(line, "soil-to-fertilizer") {
			seedMap.Mode = MAP_MODE_FERT
		} else if strings.HasPrefix(line, "fertilizer-to-water") {
			seedMap.Mode = MAP_MODE_WATER
		} else if strings.HasPrefix(line, "water-to-light") {
			seedMap.Mode = MAP_MODE_LIGHT
		} else if strings.HasPrefix(line, "light-to-temperature") {
			seedMap.Mode = MAP_MODE_TEMP
		} else if strings.HasPrefix(line, "temperature-to-humidity") {
			seedMap.Mode = MAP_MODE_HUMIDITY
		} else if strings.HasPrefix(line, "humidity-to-location") {
			seedMap.Mode = MAP_MODE_LOCATION
		} else if len(line) > 0 {
			seedMap.parseMapLine(line)
		}
	}

	return seedMap
}

func part1(lines []string) int {
	seedMap := NewSeedMap(lines, true)

	var lowestLoc = -1
	for _, s := range seedMap.Seeds {
		values := seedMap.getSeedValues(s)
		if lowestLoc == -1 || values.Location < lowestLoc {
			lowestLoc = values.Location
		}
	}

	return lowestLoc
}

func part2(lines []string) int {
	seedMap := NewSeedMap(lines, false)

	for i := 0; i < math.MaxInt; i++ {
		seed := seedMap.getSeedFromLoc(i)
		if seed != -1 {
			return i
		}
	}

	return -1
}

func main() {
	file, scanner := input.OpenInputText("./input/day5.txt")
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
