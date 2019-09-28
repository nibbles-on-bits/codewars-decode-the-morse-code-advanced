package main

import (
	"fmt"
	"strings"
)


func main() {
	test := "1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"
	fmt.Println(getWordsFromBitSentence(test))
}




//DecodeBits will take a string in the form of 1's and 0's and
// return the morse code equivilant in .'s and -'s
func DecodeBits(bits string) string {
	ret := ""



	return ret
}

//DecodeMorse will take a Morse encoded sentence in the form of
// .'s and -'s and return the alphabetic version.
func DecodeMorse(morseCode string) string {
	ret := ""

	// get timing
	// shrink
	// seperate into word
	//  seperate into letters
	//   translate letters from 1's and 0's to .'s and dashes
	// reconstruct

	timing := GetTimingFromBitSentence(morseCode)
	shrunk := ShrinkBitSentence(morseCode, timing)		// some translators are slower than others
	bitWords := strings.Split(shrunk, "0000000")				// break apart the words

	for _, bitWord := range bitWords {
		fmt.Printf("mcWord = %v\n", bitWord)
		bitLetters := strings.Split(bitWord, "000")
		for _, bitLetter := range bitLetters {
			morseLetter := BitLetterToMorseCodeLetter(bitLetter)
			fmt.Printf("morseLetter = %v\n", morseLetter)
			ret += morseLetter
			//ret += string(MorseCodeLookupChar(morseLetter))
		}
		ret += " "
		
		fmt.Println(bitLetters)
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



func getWordsFromBitSentence(bitSentence string) []string {
	ret := []string{}
		
	timing := GetTimingFromBitSentence(bitSentence)
	shrunk := ShrinkBitSentence(bitSentence, timing)		// some translators are slower than others
	bitWords := strings.Split(shrunk, "0000000")				// break apart the words

	for _, bitWord := range bitWords {
		fmt.Printf("mcWord = %v\n", bitWord)
		bitLetters := strings.Split(bitWord, "000")
		for _, bitLetter := range bitLetters {
			morseLetter := BitLetterToMorseCodeLetter(bitLetter)
			fmt.Printf("morseLetter = %v\n", morseLetter)
			//ret += string(MorseCodeLookupChar(morseLetter))
		}
		
		fmt.Println(bitLetters)
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

// DiscoverTimingFromBitSentence will determine the number of characters that represents
// a single unit of time.
func GetTimingFromBitSentence(bitSentence string) int {
	ret := 9999
	lastChar := '0'
	count := 0
	
	for _, v := range bitSentence {
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