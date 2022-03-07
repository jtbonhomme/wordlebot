package wordle_test

import (
	"testing"

	"github.com/jtbonhomme/wordlebot/internal/wordle"
)

func TestTry(t *testing.T) {
	type TestCase struct {
		word           string
		guess          string
		expectedResult string
		expectedError  error
	}
	type TestCases []TestCase

	tests := TestCases{
		TestCase{
			word:           "abcde",
			guess:          "abcde",
			expectedResult: "22222",
			expectedError:  nil,
		},
		TestCase{
			word:           "tapis",
			guess:          "tatin",
			expectedResult: "22020",
			expectedError:  nil,
		},
		TestCase{
			word:           "tarte",
			guess:          "tatin",
			expectedResult: "22100",
			expectedError:  nil,
		},
		TestCase{
			word:           "tarte",
			guess:          "artin",
			expectedResult: "11100",
			expectedError:  nil,
		},
		TestCase{
			word:           "tarte",
			guess:          "attin",
			expectedResult: "11100",
			expectedError:  nil,
		},
		TestCase{
			word:           "abcde",
			guess:          "awxyz",
			expectedResult: "20000",
			expectedError:  nil,
		},
		TestCase{
			word:           "abcde",
			guess:          "aexyz",
			expectedResult: "21000",
			expectedError:  nil,
		},
		TestCase{
			word:           "abcde",
			guess:          "aecyz",
			expectedResult: "21200",
			expectedError:  nil,
		},
		TestCase{
			word:           "abcde",
			guess:          "aecbz",
			expectedResult: "21210",
			expectedError:  nil,
		},
		TestCase{
			word:           "abcde",
			guess:          "aecbd",
			expectedResult: "21211",
			expectedError:  nil,
		},
	}

	for _, test := range tests {
		got, err := wordle.Try(test.word, test.guess, false)
		if err != nil && test.expectedError == nil {
			t.Errorf("unexpected error with word %s and guess %s: got %w (expected: nil)", test.word, test.guess, err)
		}
		if got != test.expectedResult {
			t.Errorf("unexpected result with word %s and guess %s: got %s (expected: %s)", test.word, test.guess, got, test.expectedResult)
		}
	}
}

func TestCountLetter(t *testing.T) {
	type TestCase struct {
		letter         rune
		word           string
		expectedResult int
	}
	type TestCases []TestCase

	tests := TestCases{
		TestCase{
			letter:         rune('z'),
			word:           "abcde",
			expectedResult: 0,
		},
		TestCase{
			letter:         rune('a'),
			word:           "abcde",
			expectedResult: 1,
		},
		TestCase{
			letter:         rune('e'),
			word:           "abcde",
			expectedResult: 1,
		},
		TestCase{
			letter:         rune('e'),
			word:           "aecde",
			expectedResult: 2,
		},
		TestCase{
			letter:         rune('a'),
			word:           "aecde",
			expectedResult: 1,
		},
		TestCase{
			letter:         rune('a'),
			word:           "aecda",
			expectedResult: 2,
		},
		TestCase{
			letter:         rune('a'),
			word:           "",
			expectedResult: 0,
		},
	}

	for _, test := range tests {
		got := wordle.CountLetter(test.letter, test.word)
		if got != test.expectedResult {
			t.Errorf("unexpected result with word %s and letter %s: got %d (expected: %d)", test.word, string(test.letter), got, test.expectedResult)
		}
	}
}
