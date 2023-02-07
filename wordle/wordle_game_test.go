package wordle

import (
	"testing"
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