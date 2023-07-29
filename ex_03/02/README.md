# Exercise 3.1: Count The Number Of Different Integers In A String.

## Requirement
- Goal: Count the number of different integers in a String.
- Outputs: Number of different integers.

Given that a string word consists of digits and lowercase English letters, 2 integers are considered different if their decimal representation without any leading zeros are different.

## Example
```bash
"a123bc34d8ef34" => 3 (123, 34, 8) 
"A1b01c001" => 1 (1)
```

```bash
$ go run main.go
3
```
