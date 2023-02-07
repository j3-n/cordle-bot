package wordle

import (
	"testing"
	"reflect"
)

func TestValidateGuess(t *testing.T){
	var tests = []struct{
		guess 	string
		outcome error
	}{
		{"abcde", 	nil},
		{"abCde", 	ErrInvalidFormat},
		{"ab.de", 	ErrInvalidFormat},
		{"abcd", 	ErrInvalidLength},
		{"abcdef", 	ErrInvalidLength},
	}

	for _, test := range tests{
		t.Run(test.guess, func(t *testing.T){
			outcome := validateGuess(test.guess)
			if outcome != test.outcome{
				t.Errorf("got '%v', want '%v'", outcome, test.outcome)
			}
		})
	}
}

func TestCountRunes(t *testing.T){
	var tests = []struct{
		input 	string
		outcome map[rune]int
	}{
		{"aaaaa", 	map[rune]int{'a': 5}},
		{"abcde", 	map[rune]int{'a': 1, 'b': 1, 'c': 1, 'd': 1, 'e': 1}},
		{"abbbb", 	map[rune]int{'a': 1, 'b': 4}},
	}

	for _, test := range tests{
		t.Run(test.input, func(t *testing.T){
			outcome := countRunes(test.input)
			if !reflect.DeepEqual(outcome, test.outcome){
				t.Errorf("got '%v', want '%v'", outcome, test.outcome)
			}
		})
	}
}