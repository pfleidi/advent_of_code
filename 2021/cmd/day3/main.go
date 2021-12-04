package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pfleidi/advent_of_code/2021/pkg/input"
)

func main() {
	readings, err := input.ReadStdInLines()
	if err != nil {
		log.Fatal(err)
	}

	result1, err := Exercise1(readings)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The exercise 1 result is %d!\n\n", result1)

	result2, err := Exercise2(readings)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The exercise 2 result is %d!\n\n", result2)
}

func Exercise1(readings []string) (int, error) {
	gammaBinary := ""
	epsilonBinary := ""

	for columnPosition := 0; columnPosition < len(readings[0]); columnPosition++ {
		mostCommonBit := getMostCommonBit(readings, columnPosition)

		if mostCommonBit == "0" {
			gammaBinary = gammaBinary + "0"
			epsilonBinary = epsilonBinary + "1"
		} else {
			gammaBinary = gammaBinary + "1"
			epsilonBinary = epsilonBinary + "0"
		}
	}

	gammaRate, err := strconv.ParseInt(gammaBinary, 2, 64)
	if err != nil {
		return 0, err
	}

	epsilonRate, err := strconv.ParseInt(epsilonBinary, 2, 64)
	if err != nil {
		return 0, err
	}

	log.Printf("gammaRate: %d, epsilonRate: %d", gammaRate, epsilonRate)

	return int(epsilonRate) * int(gammaRate), nil
}

func Exercise2(readings []string) (int, error) {
	oxygenBinary := scanBits(readings, 0, "", getMostCommonBit)
	oxygenReading, err := strconv.ParseInt(oxygenBinary, 2, 64)
	if err != nil {
		return 0, err
	}

	co2ScrubberBinary := scanBits(readings, 0, "", getLeastCommonBits)
	co2ScrubberReading, err := strconv.ParseInt(co2ScrubberBinary, 2, 64)
	if err != nil {
		return 0, err
	}

	log.Printf("oxygenReading: %d, co2ScrubberReading: %d\n", oxygenReading, co2ScrubberReading)

	return int(oxygenReading) * int(co2ScrubberReading), nil
}

func scanBits(readings []string, columnPosition int, searchPrefix string, getBit func([]string, int) string) string {
	if len(readings) == 1 {
		return readings[0]
	}

	newPrefix := searchPrefix + getBit(readings, columnPosition)
	filteredReadings := filterReadingsByPrefix(readings, newPrefix)

	return scanBits(filteredReadings, columnPosition+1, newPrefix, getBit)
}

func getMostCommonBit(readings []string, position int) string {
	bitCounts := make(map[string]int)

	for _, reading := range readings {
		currentBit := string(reading[position])
		bitCounts[currentBit] = bitCounts[currentBit] + 1
	}

	if bitCounts["0"] > bitCounts["1"] {
		return "0"
	} else {
		return "1"
	}
}

func getLeastCommonBits(readings []string, position int) string {
	bitCounts := make(map[string]int)

	for _, reading := range readings {
		currentBit := string(reading[position])
		bitCounts[currentBit] = bitCounts[currentBit] + 1
	}

	if bitCounts["0"] > bitCounts["1"] {
		return "1"
	} else {
		return "0"
	}
}

func filterReadingsByPrefix(readings []string, prefix string) []string {
	results := make([]string, 0)

	for _, reading := range readings {
		if strings.HasPrefix(reading, prefix) {
			results = append(results, reading)
		}
	}

	return results
}
