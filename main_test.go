package main

import (
	"testing"
	"fmt"
)


func TestGetWordsFromBitSentence() {

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

func TestBitStringToMorseCodeString(t *testing.T) {
	type TestCase struct {
		onesAndZeros string
		dotsAndDashes string
	}

	testCases := []TestCase {
		{"10101010", "...."},
		{"11101010", "-..."},
	}

	for _, tc := range testCases {
		got := BitStringToMorseCodeString(tc.onesAndZeros)
		want := tc.dotsAndDashes

		if got != want {
			t.Errorf("got: %v\nwant: %v", got, want)
		}
	}


	got := BitStringToMorseCodeString("1010101")
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



