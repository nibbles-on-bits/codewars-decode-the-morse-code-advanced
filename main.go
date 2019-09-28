package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("codewars-decode-the-morse-code-advanced.  Please run unit test")
}

//DecodeMorse will take a Morse encoded sentence in the form of
// .'s and -'s and return the alphabetic version.
func DecodeMorse(morseCode string) string {
	ret := ""

	words := strings.Split(morseCode, "  ")
  
	for _, word := range words {
		letters := strings.Split(word, " ")
		for _, char := range letters {
			ret += string(MorseCodeLookupChar(char))
		}
    ret += " "
	}


	return strings.Trim(ret," ")
}

//DecodeBits will take a string in the form of 1's and 0's and
// return the morse code equivilant in .'s and -'s
func DecodeBits(bits string) string {
	ret := ""

	// get timing
	// shrink
	// seperate into word
	//  seperate into letters
	//   translate letters from 1's and 0's to .'s and dashes
	// reconstruct

	timing := GetTimingFromBitSentence(bits)
	shrunk := ShrinkBitSentence(bits, timing)		// some translators are slower than others
	bitWords := strings.Split(shrunk, "0000000")		// break apart the words

	for _, bitWord := range bitWords {
		bitLetters := strings.Split(bitWord, "000")
		for _, bitLetter := range bitLetters {
			morseLetter := BitLetterToMorseCodeLetter(bitLetter)
			ret += morseLetter + " "
		}
		ret += " "
	}

	ret = strings.Trim(ret," ")
	return ret
}

func ShrinkBitSentence(bitSentence string, timing int) string {
	ret := ""
	for i, v := range bitSentence {
		if (i % timing == 0) {
			ret += string(v)
		}
	}
	return ret
}


// BitLetterToMorseCodeLetter will take a string in the form of "10101010" and convert to morse code "...."
//  or 10111 to ".-" 
func BitLetterToMorseCodeLetter(bitSequence string) string {
	ret := ""

	arr := strings.Split(bitSequence,"0")
	for _, c := range arr {
		if c=="1" { ret+= "."}	
		if c=="111" { ret+= "-"}	
	}

	return ret
}

// GetTimingFromBitSentence will determine the number of characters that represents
// a single unit of time.
func GetTimingFromBitSentence(bitSentence string) int {

	bitSentence = strings.Trim(bitSentence, "0")
	// Strip leading and trailing 0's off
	// Find the smallest grouping of 0's
	// Find the smallest grouping of 1's
	// return the smaller of the two
	if len(bitSentence)<=2 { return 1}
	if bitSentence[0:2] == "10" {
		return 1 
   }


	lastChar := ' '
	smallestGroup0s := 99999
	smallestGroup1s := 99999


	g0 := 0
	g1 := 0

	for _, thisChar := range bitSentence {

		if thisChar=='0' { g0++ }
		if thisChar=='1' { g1++ }

		if thisChar == '0' && lastChar == '1' {
			if g1 < smallestGroup1s {
				smallestGroup1s = g1
			}	
			g1 = 0	//reset
		}

		if thisChar == '1' && lastChar == '0' {
			if g0 < smallestGroup0s {
				smallestGroup0s = g0
			}
			g0 = 0 // reset
		}

		lastChar = thisChar

	}

	if (smallestGroup1s == 99999) {
		smallestGroup1s = g1
	}
	
	if smallestGroup0s <= smallestGroup1s {
		return smallestGroup0s
	}

	if smallestGroup1s < smallestGroup0s {
		return smallestGroup1s
	}

	return 1
}

func MorseCodeLookupChar(seq string) byte {
	ret := byte(' ')

	m := map[string]byte {
		".-"	:'A',
		"-..."	:'B',
		"-.-."	:'C',
		"-.."	:'D',
		"."		:'E',
		"..-."	:'F',
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
		".-.-.-":'.',
	}

	ret = m[seq]

	return ret
}