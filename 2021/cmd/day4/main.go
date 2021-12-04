package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/pfleidi/advent_of_code/2021/pkg/input"
)

var whiteSpace = regexp.MustCompile(`\s+`)

func main() {
	boards, calledNumbers, err := readBoardInput()
	if err != nil {
		log.Fatal(err)
	}

	result1, err := Exercise1(boards, calledNumbers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The exercise 1 result is %d!\n\n", result1)

	result2, err := Exercise2(boards, calledNumbers)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The exercise 2 result is %d!\n\n", result2)

}

func Exercise1(boards []board, numbers []int) (int, error) {
	winningBoard, lastNumber, err := findWinningBoard(boards, numbers)

	if err != nil {
		return 0, err
	}

	fmt.Println("Found winning board at number:", lastNumber)
	winningBoard.print()
	score := winningBoard.getScore(lastNumber)

	return score, nil
}

func Exercise2(boards []board, numbers []int) (int, error) {
	winningBoard, lastNumber, err := findLastWinningBoard(boards, numbers)

	if err != nil {
		return 0, err
	}

	fmt.Println("Found last winning board at number:", lastNumber)
	winningBoard.print()
	score := winningBoard.getScore(lastNumber)

	return score, nil
}

func findWinningBoard(boards []board, numbers []int) (board, int, error) {
	for _, number := range numbers {
		for _, board := range boards {
			board.markNumber(number)

			if board.isWinner() {
				return board, number, nil
			}
		}
	}

	return board{}, 0, fmt.Errorf("could not find a winning board")
}

func findLastWinningBoard(boards []board, numbers []int) (board, int, error) {
	winningBoardIndexes := make([]int, 0)

	for _, number := range numbers {
		for boardIndex, board := range boards {
			if isBoardSkipped(winningBoardIndexes, boardIndex) {
				continue
			}

			board.markNumber(number)

			if board.isWinner() {
				if len(boards)-len(winningBoardIndexes) == 1 {
					return board, number, nil
				}

				winningBoardIndexes = append(winningBoardIndexes, boardIndex)
			}
		}
	}

	return board{}, 0, fmt.Errorf("could not find a winning board")
}

func isBoardSkipped(winningBoardIndexes []int, boardIndex int) bool {
	for _, winningBoard := range winningBoardIndexes {
		if boardIndex == winningBoard {
			return true
		}
	}

	return false

}

type coordinate struct {
	number int
	marked bool
}

func (c *coordinate) String() string {
	if c.marked {
		return fmt.Sprintf("*%d*", c.number)
	} else {
		return fmt.Sprintf("%d", c.number)
	}
}

type board struct {
	rows [][]*coordinate
}

func (b board) markNumber(number int) {
	for _, row := range b.rows {
		for _, coordinate := range row {
			if coordinate.number == number {
				coordinate.marked = true
			}
		}
	}
}

func (b board) isWinner() bool {
	// first check all rows
	for _, row := range b.rows {
		for columnIndex, coordinate := range row {
			if !coordinate.marked {
				break
			}

			if columnIndex == len(row)-1 {
				return true
			}
		}
	}

	// then check all columns
	for columnIndex := 0; columnIndex < len(b.rows[0]); columnIndex++ {
		for rowIndex, row := range b.rows {
			if !row[columnIndex].marked {
				break
			}

			if rowIndex == len(b.rows)-1 {
				return true
			}
		}
	}

	return false
}

func (b board) getScore(lastNumber int) int {
	unmarkedNumbers := 0

	for _, row := range b.rows {
		for _, coordinate := range row {
			if !coordinate.marked {
				unmarkedNumbers = unmarkedNumbers + coordinate.number
			}
		}
	}

	return unmarkedNumbers * lastNumber
}

func (b board) print() {
	for _, row := range b.rows {
		rowStrings := make([]string, len(row))

		for i, coordinate := range row {
			rowStrings[i] = coordinate.String()
		}

		fmt.Println(strings.Join(rowStrings, "\t"))
	}
	fmt.Println("")
}

func readBoardInput() ([]board, []int, error) {
	input, err := input.ReadStdInLines()
	if err != nil {
		return nil, nil, err
	}

	numberStrings := strings.Split(input[0], ",")
	calledNumbers, err := stringsToNumbers(numberStrings)
	if err != nil {
		return nil, nil, err
	}

	boardStrings := input[2:]
	boards := make([]board, 0)
	rows := make([][]*coordinate, 0)

	for index, line := range boardStrings {
		if line == "" {
			continue
		}

		rowNumberStrings := whiteSpace.Split(line, -1)
		row, err := stringsToCoordinates(rowNumberStrings)
		if err != nil {
			return nil, nil, err
		}

		rows = append(rows, row)

		if index == len(boardStrings)-1 || boardStrings[index+1] == "" {
			boards = append(boards, board{rows: rows})
			rows = make([][]*coordinate, 0)
		}

	}

	return boards, calledNumbers, nil
}

func stringsToNumbers(numberStrings []string) ([]int, error) {
	numbers := make([]int, len(numberStrings))

	for i, stringNumber := range numberStrings {
		number, err := strconv.Atoi(stringNumber)
		if err != nil {
			return nil, err
		}

		numbers[i] = number
	}

	return numbers, nil
}

func stringsToCoordinates(numberStrings []string) ([]*coordinate, error) {
	coordinates := make([]*coordinate, 0)

	for _, stringNumber := range numberStrings {
		if stringNumber == "" {
			continue
		}
		number, err := strconv.Atoi(stringNumber)
		if err != nil {
			return nil, err
		}

		coordinates = append(coordinates, &coordinate{number: number})
	}

	return coordinates, nil
}
