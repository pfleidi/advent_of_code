package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"

	"github.com/pfleidi/advent_of_code/2021/pkg/input"
)

var inputLinePattern = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

func main() {
	input, err := input.ReadStdInLines()
	if err != nil {
		log.Fatal(err)
	}

	lines, err := parseLines(input)
	if err != nil {
		log.Fatal(err)
	}

	result1, err := Exercise1(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The exercise 1 result is %d!\n\n", result1)

	result2, err := Exercise2(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The exercise 2 result is %d!\n\n", result2)

}

type point struct {
	x int
	y int
}

func (p point) pathTo(p2 point, diagonal bool) ([]point, error) {
	path := make([]point, 0)

	// default steps to 0
	xStep := 0
	yStep := 0

	xDiff := p.x - p2.x
	yDiff := p.y - p2.y

	if (xDiff != 0 && yDiff != 0) && !diagonal {
		return nil, nil
	}

	if xDiff != 0 && yDiff != 0 && math.Abs(float64(xDiff)) != math.Abs(float64(yDiff)) {
		return nil, fmt.Errorf("can only move in horizontal, vertical, or in 45º steps")
	}

	if xDiff > 0 {
		xStep = -1
	} else if xDiff < 0 {
		xStep = 1
	}

	if yDiff > 0 {
		yStep = -1
	} else if yDiff < 0 {
		yStep = 1
	}

	stepCount := int(math.Max(math.Abs(float64(xDiff)), math.Abs(float64(yDiff))))

	for i := 0; i <= stepCount; i++ {
		pointX := p.x + xStep*i
		pointY := p.y + yStep*i
		path = append(path, point{x: pointX, y: pointY})
	}

	return path, nil
}

func (p point) getMaxCoordinates(p2 point) (int, int) {
	maxX := 0
	maxY := 0

	if p.x > p2.x {
		maxX = p.x
	} else {
		maxX = p2.x
	}

	if p.y > p2.y {
		maxY = p.y
	} else {
		maxY = p2.y
	}

	return maxX, maxY
}

type line struct {
	start point
	end   point
}

type grid struct {
	lines       []line
	coordinates [][]int
}

func newGrid(lines []line) *grid {
	xSize, ySize := getGridSize(lines)
	coordinates := make([][]int, xSize)

	for i := 0; i < xSize; i++ {
		coordinates[i] = make([]int, ySize)
	}

	return &grid{lines: lines, coordinates: coordinates}
}

func (g *grid) print() {
	if os.Getenv("DEBUG") != "true" {
		return
	}

	fmt.Print("  ")
	for x := 0; x < len(g.coordinates[0]); x++ {
		fmt.Printf("\033[31m%-2d", x)
	}

	fmt.Print("\033[39m\n")

	for y := 0; y < len(g.coordinates); y++ {
		fmt.Printf("\033[31m%-2d\033[39m", y)

		for x := 0; x < len(g.coordinates[y]); x++ {
			count := g.coordinates[x][y]

			if count == 0 {
				fmt.Printf("%-2s", ".")
			} else {
				fmt.Printf("%-2d", count)
			}
		}
		fmt.Print("\n")
	}
}

func (g *grid) placeLines(diagonal bool) {
	for _, line := range g.lines {
		path, err := line.start.pathTo(line.end, diagonal)
		if err != nil && diagonal {
			fmt.Printf("Could not place line %v due to: %v\n", line, err)
		}

		if path == nil {
			continue
		}

		for _, point := range path {
			g.coordinates[point.x][point.y] = g.coordinates[point.x][point.y] + 1
		}
	}
}

func (g *grid) getOverlapCount() int {
	overlapCount := 0

	for y := 0; y < len(g.coordinates); y++ {
		for x := 0; x < len(g.coordinates[y]); x++ {
			count := g.coordinates[x][y]

			if count > 1 {
				overlapCount = overlapCount + 1
			}
		}
	}

	return overlapCount
}

func getGridSize(lines []line) (int, int) {
	maxX := 0
	maxY := 0

	for _, line := range lines {
		x, y := line.start.getMaxCoordinates(line.end)
		if x > maxX {
			maxX = x
		}

		if y > maxY {
			maxY = y
		}
	}

	return maxX + 1, maxY + 1
}

func Exercise1(lines []line) (int, error) {
	grid := newGrid(lines)

	grid.placeLines(false)
	grid.print()

	return grid.getOverlapCount(), nil
}

func Exercise2(lines []line) (int, error) {
	grid := newGrid(lines)

	grid.placeLines(true)
	grid.print()

	return grid.getOverlapCount(), nil
}

func parseLines(input []string) ([]line, error) {
	lines := make([]line, 0)

	for _, inpuLine := range input {
		match := inputLinePattern.FindStringSubmatch(inpuLine)
		startX, _ := strconv.Atoi(match[1])
		startY, _ := strconv.Atoi(match[2])
		endX, _ := strconv.Atoi(match[3])
		endY, _ := strconv.Atoi(match[4])

		line := line{
			start: point{x: startX, y: startY},
			end:   point{x: endX, y: endY},
		}

		lines = append(lines, line)
	}

	return lines, nil
}
