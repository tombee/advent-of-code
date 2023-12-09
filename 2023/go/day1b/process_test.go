package main

import (
	"reflect"
	"testing"
)

func TestProcessLine(t *testing.T) {
	type test struct {
		input []byte
		want  int
	}

	tests := []test{
		{input: []byte("two1nine"), want: 29},
		{input: []byte("eightwothree"), want: 83},
		{input: []byte("abcone2threexyz"), want: 13},
		{input: []byte("xtwone3four"), want: 24},
		{input: []byte("4nineeightseven2"), want: 42},
		{input: []byte("zoneight234"), want: 14},
		{input: []byte("7pqrstsixteen"), want: 76},
	}

	for _, tc := range tests {
		got, err := ProcessLine(tc.input)
		if err != nil {
			t.Fatalf("error: %v", err)
		}
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
