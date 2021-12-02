package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pfleidi/advent_of_code/2021/pkg/input"
)

type command struct {
	name  string
	value int
}

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
	horizontalPosition := 0
	depth := 0

	commands, err := parseReadings(readings)
	if err != nil {
		return 0, err
	}

	for _, command := range commands {
		switch command.name {
		case "forward":
			horizontalPosition = horizontalPosition + command.value
		case "up":
			depth = depth - command.value
		case "down":
			depth = depth + command.value
		default:
			return 0, fmt.Errorf("invalid command: %s", command.name)
		}

	}

	return horizontalPosition * depth, nil
}

func Exercise2(readings []string) (int, error) {
	horizontalPosition := 0
	depth := 0
	aim := 0

	commands, err := parseReadings(readings)
	if err != nil {
		return 0, err
	}

	for _, command := range commands {
		switch command.name {
		case "forward":
			horizontalPosition = horizontalPosition + command.value
			depth = depth + aim*command.value
		case "up":
			aim = aim - command.value
		case "down":
			aim = aim + command.value
		default:
			return 0, fmt.Errorf("invalid command: %s", command.name)
		}
	}

	return horizontalPosition * depth, nil
}

func parseReadings(readings []string) ([]command, error) {
	commands := make([]command, len(readings))

	for index, line := range readings {
		splits := strings.Split(line, " ")

		if len(splits) != 2 {
			return []command{}, fmt.Errorf("could not parse line '%s'", line)
		}

		commandName := splits[0]
		value, err := strconv.Atoi(splits[1])

		if err != nil {
			return []command{}, err
		}

		commands[index] = command{name: commandName, value: value}

	}

	return commands, nil
}
