package main

import (
	"fmt"
	"strings"
)


func main() {
	test := "1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"
	fmt.Println(DiscoverTiming("1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"))
	_ = getWordsFromMorseCode(test, 2)
}


func getWordsFromMorseCode(bitSequence string, timing int) []string {
	ret := []string{}
		
	// first let's shrink
	shrunk := ""	
	for i, v := range bitSequence {
		if (i % timing == 0) {
			shrunk += string(v)
		}
	}

	fmt.Printf("Shrunk : %v\n", shrunk)

	mcWords := strings.Split(shrunk, "000")
	//fmt.Println(mcWords)

	for _, mcWord := range mcWords {
		fmt.Printf("mcWord = %v\n", mcWord)
		// this is in the form of a character string of 1's and 0's ie : 1010101 (which would be .... H)

	}

	return ret
}

// BitStringToMorseCodeString will take a string in the form of "1010101" and convert to morse code "...."
//  or 10111 to ".-" 
func BitStringToMorseCodeString(bitSequence string) string {
	//

}

// DiscoverTiming will determine the number of characters that represents
// a single unit of time.
func DiscoverTiming(bitSequence string) int {
	ret := 9999
	lastChar := '0'
	count := 0
	
	for _, v := range bitSequence {
		if (v=='1') {
			count++
			lastChar = '1'
		} else if (v == '0') {
			if (lastChar == '1') {
				if count < ret { ret = count }	
			}
			lastChar = '0'
			count = 0
		}
		
	}

	return ret
}

func MorseCodeLookupChar(seq string) byte {
	ret := byte(' ')

	m := map[string]byte {
		".-"	:'A',
		"-..."	:'B',
		"-.-."	:'C',
		"-.."	:'D',
		"."		:'E',
		"..=."	:'F',
		"--."	:'G',
		"...."	:'H',
		".."	:'I',
		".---"	:'J',
		"-.-"	:'K',
		".-.."	:'L',
		"--"	:'M',
		"-."	:'N',
		"---"	:'O',
		".--."	:'P',
		"--.-"	:'Q',
		".-."	:'R',
		"..."	:'S',
		"-"		:'T',
		"..-"	:'U',
		"...-"	:'V',
		".--"	:'W',
		"-..-"	:'X',
		"-.--"	:'Y',
		"--.."	:'Z',
	}

	ret = m[seq]
	//fmt.Println()

	return ret
}