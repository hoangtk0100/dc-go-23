# Exercise 3.1: Count Rectangles.

## Requirement
- Goal: Find and count number of rectangles in a 2D array.
- Inputs: An array filled with 0s and 1s.
- Outputs: Number of rectangles filled with 1s.

Given that rectangles are separated and do not touch each other but they can touch the boundary of the array. A single element rectangle counts.

## Example
```bash
    {1, 0, 0, 0, 0, 0, 0}
    {0, 0, 0, 0, 0, 0, 0}
    {1, 0, 0, 1, 1, 1, 0}
    {0, 1, 0, 1, 1, 1, 0}
    {0, 1, 0, 0, 0, 0, 0}
    {0, 1, 0, 1, 1, 0, 0}
    {0, 0, 0, 1, 1, 0, 0}
    {0, 0, 0, 0, 0, 0, 1}
```

```bash
$ go run main.go
6
```
