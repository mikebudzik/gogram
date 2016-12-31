package main

import (
	"fmt"
)

// return the number of occurrences of letter in gram.  Just in case they pass a multi-char string for letter, just use the first char in letter.
func count(gram string, letter string) (count int) {
	var iLength int = len(gram)
	count = 0
	for i := 0; i < iLength; i++ {
		if gram[i:i+1] == letter[0:1] {
			count++
		}
	}
	
	return
}

func main() {
	// todo
	// learn how to make an array to hold counts
	// compare counts to most letter frequency (E and T are most common in general)
	// Look for two letter words 
	//   TO is common and T is common - look for 2 letter words with first letter having high occurrance rate
	//   SO is similar to TO - be sure to consider that - perhaps the O is safe
	//   AS has a common 2nd letter
	// look for single letter words, they are probably A, I, or possibly O
	// think about apostraphies I'M or XXXXXX'S

	var gram string = "BWIJLIC BVOD TVD WI KIC DX MWQ EKUXQK MWQ TVDO WI KIC DX BWIJLIC."
	fmt.Println("The crpytogram is \n\n"+gram)
	for i := 0; i < 26; i++ {
		// loop over the letters of the alphabet
		fmt.Printf("%q: %d\n", i+65, count(gram, string(i+65)))
	}
}

