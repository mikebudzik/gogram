package main

import (
	"fmt"
	"strings"
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
	// print a line for the user to look at possible solutions
	// compare counts to most letter frequency (E and T are most common in general)
	// Look for two letter words 
	//   TO is common and T is common - look for 2 letter words with first letter having high occurrance rate
	//   SO is similar to TO - be sure to consider that - perhaps the O is safe
	//   AS has a common 2nd letter
	// look for single letter words, they are probably A, I, or possibly O
	// think about apostraphies I'M or XXXXXX'S

	var gram string = "KIF HR WBF RIIA FPI JINHYKUMOW OWRVIJ BJ FPI GIDBMJOFUM OWRVIJ, YHF FPI JUZPF OWRVIJ. KIF HR WBF RIIA FB TUC FPI YKODI TBJ FPI NORF. KIF HR OMMINF BHJ BVW JIRNBWRUYUKUFS TBJ FPI THFHJI."
	var counts [26]int
	var sFreq1Char string = ""
	var iFreq1Count = 0
	var sFreq2Char string = ""
	var iFreq2Count = 0
	
	for i := 0; i < 26; i++ {
		// loop over the letters of the alphabet
		counts[i] = count(gram, string(i+65))
		//fmt.Printf("%q: %d\n", i+65, counts[i])
		
		if counts[i] == iFreq1Count {
			if iFreq1Count == iFreq2Count {
				iFreq2Count = counts[i]
				sFreq2Char = string(i+65)
			} else {
				iFreq2Count = iFreq1Count 
				sFreq2Char = sFreq1Char 
				iFreq1Count = counts[i]
				sFreq1Char = string(i+65)				
			}
		} else if counts[i] > iFreq1Count{
				iFreq2Count = iFreq1Count 
				sFreq2Char = sFreq1Char 
				iFreq1Count = counts[i]
				sFreq1Char = string(i+65)	
		} else if counts[i] > iFreq2Count {
			iFreq2Count = counts[i]
			sFreq2Char = string(i+65)
		}
	}
	
	fmt.Printf("%s: %d\n", sFreq1Char, iFreq1Count)
	fmt.Printf("%s: %d\n", sFreq2Char, iFreq2Count)
		
	// break it up into words
	fmt.Println("The crpytogram is \n")
	var words []string = strings.Split(gram, " ")
	for i := 0; i < len(words); i++ {
		fmt.Printf("%s ", words[i])
	}
	
	fmt.Println()
	
	// print a line with blanks for guesses
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ { 
			fmt.Printf("_")
		}
		fmt.Printf(" ")
	}
	
	fmt.Println()
	
	// print a line with blanks for guesses, but assume E then T are the most common in the input
	fmt.Println("Guess 1")
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ { 
			if words[i][j:j+1] == sFreq1Char {
				fmt.Printf("E")
			} else if words[i][j:j+1] == sFreq2Char {
				fmt.Printf("T")
			} else {
				fmt.Printf("_")
			}
		}
		fmt.Printf(" ")
	}
	
	// print a line with blanks for guesses, but assume T then E are the most common in the input
	fmt.Println("Guess 2")
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ { 
			if words[i][j:j+1] == sFreq1Char {
				fmt.Printf("T")
			} else if words[i][j:j+1] == sFreq2Char {
				fmt.Printf("E")
			} else {
				fmt.Printf("_")
			}
		}
		fmt.Printf(" ")
	}
}
