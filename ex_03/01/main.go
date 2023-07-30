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

	// markPoints marks visited points in detected rectangle
	var markPoints func(row, col int)
	markPoints = func(row, col int) {
		if row < 0 || row >= rows || col < 0 || col >= cols || rectangles[row][col] != 1 {
			return
		}

		rectangles[row][col] = -1

		markPoints(row-1, col)
		markPoints(row+1, col)
		markPoints(row, col-1)
		markPoints(row, col+1)
	}

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if rectangles[row][col] == 1 {
				count++
				markPoints(row, col)
			}
		}
	}

	return count
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
