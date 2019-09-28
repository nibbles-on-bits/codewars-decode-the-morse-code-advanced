package main

import (
	"testing"
)

//MorseCodeLookupChar

type TestCase struct {
	seq	string
	char byte
}


func TestMorseCodeLookupChar(t *testing.T) {
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



