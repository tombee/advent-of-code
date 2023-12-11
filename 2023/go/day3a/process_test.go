package main

import (
	"reflect"
	"testing"
)

func TestGetAdjacentIndicesStartLine(t *testing.T) {
	got := GetAdjacentIndices(0, 0, 10, 10, 3)
	want := [][]int{
		{0, 1},
		{1, 1},
		{2, 1},
		{3, 1},
		{3, 0},
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetAdjacentIndicesLastLine(t *testing.T) {
	got := GetAdjacentIndices(1, 1, 4, 2, 3)
	want := [][]int{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{0, 1},
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetAdjacentIndicesAtEnd(t *testing.T) {
	got := GetAdjacentIndices(1, 1, 4, 3, 3)
	want := [][]int{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{0, 2},
		{1, 2},
		{2, 2},
		{3, 2},
		{0, 1},
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetAdjacentIndicesAtEdge(t *testing.T) {
	got := GetAdjacentIndices(0, 0, 4, 4, 3)
	want := [][]int{
		{0, 1},
		{1, 1},
		{2, 1},
		{3, 1},
		{3, 0},
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetAdjacentIndices(t *testing.T) {
	got := GetAdjacentIndices(1, 1, 99, 99, 3)
	want := [][]int{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 0},
		{0, 2},
		{1, 2},
		{2, 2},
		{3, 2},
		{4, 2},
		{0, 1},
		{4, 1},
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestCollectNumbers(t *testing.T) {
	input := []byte("467..114..")
	length, number := CollectNumbers(input)
	wantLength := 3
	wantNumber := 467

	if length != wantLength {
		t.Errorf("Expected length was incorrect, got %d expected %d", length, wantLength)
	}
	if number != wantNumber {
		t.Errorf("Expected number was incorrect, got %d expected %d", number, wantNumber)
	}
}

func TestProcess(t *testing.T) {
	input := []byte(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`)

	result, err := Process(input)
	want := 4361

	if err != nil {
		t.Fatalf("Error occurred when processing, %v", err)
	}
	if result != want {
		t.Errorf("Expected result was incorrect, got %d expected %d.", result, want)
	}
}
