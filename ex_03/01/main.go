package main

import (
	"fmt"
)

// countRectangles returns a number of rectangles filled with 1
func countRectangles(rectangles [][]int) int {
	if len(rectangles) == 0 || len(rectangles[0]) == 0 {
		return 0
	}

	count := 0
	rows := len(rectangles)
	cols := len(rectangles[0])
	rectangleMarkedPoints := map[string]bool{}

	// isMarkedPoint checks if a point is already marked
	isMarkedPoint := func(x, y int) bool {
		_, ok := rectangleMarkedPoints[formatKey(x, y)]
		return ok
	}

	// markPoints marks valid points in detected rectangle
	markPoints := func(x, y int) {
		for row := x; row < rows; row++ {
			if rectangles[row][y] == 0 {
				return
			}

			for col := y; col < cols; col++ {
				if rectangles[row][col] == 0 {
					break
				}
				if isMarkedPoint(row, col) {
					continue
				}

				rectangleMarkedPoints[formatKey(row, col)] = true
			}
		}
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if rectangles[row][col] == 1 && !isMarkedPoint(row, col) {
				count++
				markPoints(row, col)
			}
		}
	}

	return count
}

// formatKey formats key for map rectangleMarkedPoints
func formatKey(x, y int) string {
	return fmt.Sprintf("(%v,%v)", x, y)
}

func main() {
	arr := [][]int{
		{1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 1, 0},
		{0, 1, 0, 1, 1, 1, 0},
		{0, 1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 0, 0},
		{0, 0, 0, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}

	count := countRectangles(arr)
	fmt.Printf("%v", count)
}
