# Exercise 01: Reordering Names Based on Country Code

## Requirement
- Goal: Create a program that reorders names based on the specified country's format.
- Inputs: First name, last name, middle name (optional), and country code.
- Outputs: Reordered name based on the country's format.

## Example
```bash
# 01
$ go run main.go John Smith VN
Output: Smith John
```

```bash
# 02
$ go run main.go Emily Rose Watson US
Output: Emily Rose Watson
```