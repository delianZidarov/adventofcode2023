package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	p := 0
	if len(args) == 1 || args[1] == "1" {
		p = 1
	} else if args[1] == "2" {
		p = 2
	}

	f, err := os.Open(args[0])

	buf := make([]byte, 5745)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	end, _ := f.Read(buf)
	b := blocks(buf[:end])
	f.Close()

	seeds, err := parseSeeds(b[0])
	if err != nil {
		fmt.Println("Seeds: ", err)
	}
	seedToSoil, err := parseInputMap(b[1])
	if err != nil {
		fmt.Println("Seed to Soil: ", err)
	}

	soilToFert, err := parseInputMap(b[2])
	if err != nil {
		fmt.Println("Soil to fert: ", err)
	}

	fertToWater, err := parseInputMap(b[3])
	if err != nil {
		fmt.Println("Fert to water: ", err)
	}

	waterToLight, err := parseInputMap(b[4])
	if err != nil {
		fmt.Println("Water to light", err)
	}

	lightToTemp, err := parseInputMap(b[5])
	if err != nil {
		fmt.Println("Light to temp", err)
	}

	tempToHumid, err := parseInputMap(b[6])
	if err != nil {
		fmt.Println("Temp to humid: ", err)
	}

	humidToLoc, err := parseInputMap(b[7])
	if err != nil {
		fmt.Println("Humid to loc: ", err)
	}

	switch p {
	case 1:
		// noting a singed int returns a negative
		// noting a unsinged int returns the max number
		// this is too big for a signed int so it needs to
		// be right shifted
		min := int(^uint(0) >> 1)
		for _, seed := range seeds {
			s := findDest(seedToSoil, seed)
			s = findDest(soilToFert, s)
			s = findDest(fertToWater, s)
			s = findDest(waterToLight, s)
			s = findDest(lightToTemp, s)
			s = findDest(tempToHumid, s)
			s = findDest(humidToLoc, s)
			if s < min {
				min = s
			}
		}
		fmt.Println("Lowest Location: ", min)
	case 2:
		// naive solution
		min := int(^uint(0) >> 1)
		for i := 0; i < len(seeds); i += 2 {
			for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
				s := findDest(seedToSoil, j)
				s = findDest(soilToFert, s)
				s = findDest(fertToWater, s)
				s = findDest(waterToLight, s)
				s = findDest(lightToTemp, s)
				s = findDest(tempToHumid, s)
				s = findDest(humidToLoc, s)
				if s < min {
					min = s
				}
			}
		}
		fmt.Println("Lowest Location: ", min)
	}
}
