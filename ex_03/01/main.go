package main

import "fmt"

var (
	rectangleMarkedPoints = map[string]bool{}
)

// countRectangles returns a number of rectangles filled with 1
func countRectangles(rectangles [][]int) int {
	count := 0

	for x := 0; x < len(rectangles); x++ {
		for y := 0; y < len(rectangles[0]); y++ {
			if rectangles[x][y] == 1 && !checkMarkedPoint(x, y) {
				count++
				markRectanglePoints(rectangles, x, y)
			}
		}
	}

	return count
}

// formatKey formats key for map rectangleMarkedPoints
func formatKey(x, y int) string {
	return fmt.Sprintf("(%v,%v)", x, y)
}

// checkMarkedPoint checks if a point is already marked
func checkMarkedPoint(x, y int) bool {
	_, ok := rectangleMarkedPoints[formatKey(x, y)]
	return ok
}

// markRectanglePoints marks valid points in detected rectangle
func markRectanglePoints(rectangles [][]int, x, y int) {
	for xIndex := x; xIndex < len(rectangles); xIndex++ {
		if rectangles[xIndex][y] == 0 {
			break
		}

		for yIndex := y; yIndex < len(rectangles[xIndex]); yIndex++ {
			if rectangles[xIndex][yIndex] == 0 {
				break
			}

			if checkMarkedPoint(xIndex, yIndex) {
				continue
			}

			rectangleMarkedPoints[formatKey(xIndex, yIndex)] = true
		}
	}
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
