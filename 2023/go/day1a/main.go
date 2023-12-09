package main

import (
	"bytes"
	"fmt"
	"os"
)

func to_digit(c byte) int {
	return int(c - '0')
}

func is_digit(c byte) bool {
	return '0' <= c && c <= '9'
}

func process_line(line []byte) (int, error) {
	first := 0
	last := 0
	for _, char := range line {
		if is_digit(char) {
			first = to_digit(char)
			break
		}
	}
	// Iterate in reverse across line
	for i := range line {
		char := line[len(line)-1-i]
		if is_digit(char) {
			last = to_digit(char)
			break
		}
	}
	return first*10 + last, nil
}

func process(data []byte) (int, error) {
	answer := 0
	lines := bytes.Split(data, []byte{'\n'})
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		line_num, err := process_line(line)
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
	answer, err := process(dat)
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
