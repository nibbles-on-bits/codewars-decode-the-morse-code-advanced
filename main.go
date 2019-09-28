package main

import (
	"fmt"
	"strings"
)


func main() {
	test := "1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"
	//fmt.Println(GetTimingFromBitSentence("1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"))
	_ = getWordsFromBitSentence(test)

	//DiscoverTiming
	//timing := DiscoverTimingFromBitSentence()
	//ShrinkBitString
	//GetWordsFromBitSentence
	//GetLettersFromWords
	//GetCharactersFromBitWord
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
			morseLetter := BitStringToMorseCodeString(bitLetter)
			fmt.Println(morseLetter)
			fmt.Println(string(MorseCodeLookupChar(morseLetter)))
		}
		
		fmt.Println(bitLetters)
		
		/*
		// this is in the form of a character string of 1's and 0's ie : 1010101 (which would be .... H)
		tmp := BitStringToMorseCodeString(bitWord)
		//ret = append(ret, MorseCodeLookupChar(tmp))
		fmt.Println(MorseCodeLookupChar(tmp))*/
	}

	return ret
}

// BitStringToMorseCodeString will take a string in the form of "10101010" and convert to morse code "...."
//  or 10111 to ".-" 
// the provided bitSequence should already be shrunk
func BitStringToMorseCodeString(bitSequence string) string {
	tmp := strings.ReplaceAll(bitSequence, "1110","-")
	tmp = strings.ReplaceAll(tmp, "10", ".")
	tmp = strings.ReplaceAll(tmp, "1", ".")
	return tmp
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