package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/pfleidi/advent_of_code/2021/pkg/input"
)

type slidingWindow struct {
	name rune
	sum  int
}

func main() {
	readings, err := readInput()
	if err != nil {
		log.Fatal(err)
	}

	readingsIncreasedCount := Exercise1(readings)
	fmt.Printf("The sliding windows increased %d times!\n\n", readingsIncreasedCount)

	windowsIncreasedCount := Exercise2(readings)
	fmt.Printf("The sliding windows increased %d times!\n\n", windowsIncreasedCount)
}

func Exercise1(readings []int) int {
	increasedCount := 0

	for index, reading := range readings {
		var change string

		if index == 0 {
			change = "N/A - no previous reading"
			fmt.Printf("%d (%s)\n", reading, change)
			continue
		}

		if reading > readings[index-1] {
			change = "increased"
			increasedCount = increasedCount + 1
		} else if reading == readings[index-1] {
			change = "no change"
		} else {
			change = "decreased"
		}

		fmt.Printf("%d (%s)\n", reading, change)
	}

	return increasedCount
}

func Exercise2(readings []int) int {
	increasedCount := 0

	windows := readSlidingWindows(readings)

	for index, window := range windows {
		var change string

		if index == 0 {
			change = "N/A - no previous sum"
			fmt.Printf("%c: %d (%s)\n", window.name, window.sum, change)
			continue
		}

		if window.sum > windows[index-1].sum {
			change = "increased"
			increasedCount = increasedCount + 1
		} else if window.sum == windows[index-1].sum {
			change = "no change"
		} else {
			change = "decreased"
		}

		fmt.Printf("%c: %d (%s)\n", window.name, window.sum, change)
	}

	return increasedCount
}

func readSlidingWindows(readings []int) []slidingWindow {
	slidingWindows := make([]slidingWindow, 0)

	for index := range readings {
		var windowSum int = 0
		windowLetter := rune('A' + index)

		if index+2 == len(readings) {
			break
		}

		for i := 0; i < 3; i++ {
			windowSum = windowSum + readings[index+i]
		}

		window := slidingWindow{
			name: windowLetter,
			sum:  windowSum,
		}
		slidingWindows = append(slidingWindows, window)

	}

	return slidingWindows
}

func readInput() ([]int, error) {
	inputLines, err := input.ReadStdInLines()
	readings := make([]int, len(inputLines))

	if err != nil {
		return nil, err
	}

	for index, line := range inputLines {
		depth, err := strconv.Atoi(line)

		if err != nil {
			break
		}

		readings[index] = depth

	}

	return readings, nil
}
