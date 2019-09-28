package main

import (
	"testing"
	"fmt"
)

func TestGetTimingFromBitSentence(t *testing.T) {
	type TestCase struct {
		bits string
		timing int
	}

	testCases := []TestCase {
		{"1110111", 1},
	}

	for _, tc := range testCases {
		got := GetTimingFromBitSentence(tc.bits)
		want := tc.timing

		if got != want {
			t.Errorf("bits: %v\ngot: %v\nwant: %v", tc.bits, got, want)
		}
	}
}

func TestDecodeMorse(t *testing.T) {
	type TestCase struct {
		morse string
		alpha string
	}

	testCases := []TestCase {
		//{".... . -.--  .--- ..- -.. .", "HEY JUDE"},
		{".", "E"},
		{"..", "I"},
		{". .", "EE"},
		{".-", "A"},

	}

	for _, tc := range testCases {
		got := DecodeMorse(tc.morse)
		want := tc.alpha

		if got != want {
			t.Errorf("morse: %v\ngot: %v\nwant: %v", tc.morse, got, want)
		}
	}

}

func TestDecodeBits(t *testing.T) {
	type TestCase struct {
		bits string
		morse string
	}

	testCases := []TestCase {
		{
			bits: "1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011", 
			morse: ".... . -.--  .--- ..- -.. .",
		},
	}

	for _, tc := range testCases {
		got := DecodeBits(tc.bits)
		want := tc.morse

		if got != want {
			t.Errorf("bits: %v\ngot: %v\nwant: %v", tc.bits, got, want)
		}
	}

}


func TestDiscoverTimingFromBitSentence(t *testing.T) {
	type TestCase struct {
		bitSentence string
		expectedTiming int
	}

	testCases := []TestCase {
		{
			bitSentence: "1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011", 
			expectedTiming: 2,
		},
	}

	for _, tc := range testCases {
		got := GetTimingFromBitSentence(tc.bitSentence)
		want := tc.expectedTiming

		if got != want {
			t.Errorf("bitSentence: %v\ngot: %v\nwant: %v", tc.bitSentence, got, want)
		}
	}
}

func TestShrinkBitSentence(t *testing.T) {
	type TestCase struct {
		bitSentence string
		timing int
		shrunkSentence string
	}
	
	testCases := []TestCase {
		{
			bitSentence: "1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011", 
			timing: 2,
			shrunkSentence: "10101010001000111010111011100000001011101110111000101011100011101010001",
		},
	}

	for _, tc := range testCases {
		got := ShrinkBitSentence(tc.bitSentence, tc.timing)
		want := tc.shrunkSentence

		if got != want {
			t.Errorf("bitSentence: %v\ngot: %v\nwant: %v", tc.bitSentence, got, want)
		}
	}

}

func TestBitLetterToMorseCodeLetter(t *testing.T) {
	type TestCase struct {
		onesAndZeros string
		dotsAndDashes string
	}

	testCases := []TestCase {
		{"1011101110111", ".---"},
		{"10101010", "...."},
		{"111010101", "-..."},
	}

	for _, tc := range testCases {
		got := BitLetterToMorseCodeLetter(tc.onesAndZeros)
		want := tc.dotsAndDashes

		if got != want {
			t.Errorf("got: %v\nwant: %v", got, want)
		}
	}


	got := BitLetterToMorseCodeLetter("1010101")
	fmt.Println(got)
}

func TestMorseCodeLookupChar(t *testing.T) {

	type TestCase struct {
		seq	string
		char byte
	}

	testCases := []TestCase{
		{".-", 'A'},
		{"-...", 'B'},
	}

	for _, tc := range testCases {
		got := MorseCodeLookupChar(tc.seq)
		want := tc.char

		if got != want {
			t.Errorf("got: %v\nwant: %v", got, want)
		}
	}
}



