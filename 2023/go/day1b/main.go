package main

import (
	"bytes"
	"fmt"
	"os"
)

var NUMS = [9][]byte{
	[]byte("one"),
	[]byte("two"),
	[]byte("three"),
	[]byte("four"),
	[]byte("five"),
	[]byte("six"),
	[]byte("seven"),
	[]byte("eight"),
	[]byte("nine"),
}

func to_digit(c byte) int {
	return int(c - '0')
}

func is_digit(c byte) bool {
	return '0' <= c && c <= '9'
}

func ProcessLine(line []byte) (int, error) {
	first := 0
	last := 0

first:
	for i := 0; i < len(line); i++ {
		if is_digit(line[i]) {
			first = to_digit(line[i])
			break
		}
		for j, num := range NUMS {
			if bytes.HasPrefix(line[i:], num) {
				first = j + 1
				break first
			}
		}
	}
	// Iterate in reverse across line
last:
	for i := 0; i < len(line); i++ {
		char := line[len(line)-1-i]
		if is_digit(char) {
			last = to_digit(char)
			break
		}
		for j, num := range NUMS {
			if bytes.HasSuffix(line[:len(line)-i], num) {
				last = j + 1
				break last
			}
		}
	}
	return first*10 + last, nil
}

func Process(data []byte) (int, error) {
	answer := 0
	lines := bytes.Split(data, []byte{'\n'})
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		line_num, err := ProcessLine(line)
		if err != nil {
			return 0, err
		}
		answer += line_num
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
