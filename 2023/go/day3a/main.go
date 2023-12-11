package main

import (
	"bytes"
	"fmt"
	"os"
)

// Example data
// ------------
// 467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..

func to_digit(c byte) int {
	return int(c - '0')
}

func is_digit(c byte) bool {
	return '0' <= c && c <= '9'
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func GetAdjacentIndices(x int, y int, maxX int, maxY int, length int) [][]int {
	result := [][]int{}
	for _, j := range []int{y - 1, y + 1} {
		if j < 0 || j >= maxY {
			continue
		}
		for i := max(0, x-1); i < maxX && i < x+length+1; i++ {
			result = append(result, []int{i, j})
		}
	}
	if x-1 >= 0 {
		result = append(result, []int{x - 1, y})
	}
	if x+length < maxX {
		result = append(result, []int{x + length, y})
	}
	return result
}

func CollectNumbers(line []byte) (length int, number int) {
	if len(line) == 0 {
		return 0, 0
	}

	length = 0
	number = 0

	for i := 0; i < len(line) && is_digit(line[i]); i++ {
		number = number*10 + to_digit(line[i])
		length++
	}

	return length, number
}

func isSymbol(b byte) bool {
	for _, s := range []byte{'#', '+', '*', '$', '/', '&', '%', '=', '@', '-'} {
		if b == s {
			return true
		}
	}
	return false
}

func Process(data []byte) (int, error) {
	answer := 0
	lines := bytes.Split(data, []byte{'\n'})

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			length, number := CollectNumbers(lines[y][x:])
			if length > 0 {
				adjacents := GetAdjacentIndices(x, y, len(lines[y]), len(lines)-1, length)
				for _, adjacent := range adjacents {
					if isSymbol(lines[adjacent[1]][adjacent[0]]) {
						answer += number
						break
					}
				}
				x += length - 1
			}
		}
	}

	return answer, nil
}

func run() error {
	dat, err := os.ReadFile("input.txt")
	if err != nil {
		return err
	}
	answer, err := Process(dat)
	if err != nil {
		return err
	}
	fmt.Printf("Answer: %d\n", answer)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
