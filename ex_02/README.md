# Exercise 02: Sort Inputs Provided By The User.

## Requirement
- Goal: Create a package and a command-line tool to sort input provided by the user.
- Inputs: Number (integer or float) array, string array.
- Outputs: Sorted result based on the provided input type.

## Example
```bash
# 01
$ go run sorter.go -int 5 2 18 1
Output: 1 2 5 18
```

```bash
# 02
$ go run sorter.go -string apple orange banana
Output: apple banana orange
```

```bash
# 03
$ go run sorter.go -mix 5.5 apple 2.7 orange 3 banana
Output: 2.7 3 5.5 apple banana orange
```