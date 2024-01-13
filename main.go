package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PrintString(Standard []string, TextArray []string) {
	var word []rune
	var LineStandard int
	var Line string
	NewLine := 0 
	for str := 0; str < len(TextArray); str++ {
		word = []rune(TextArray[str])
		if len(word) > 0 {
			for Linenum := 1; Linenum <= 8; Linenum++ {
				for chr := 0; chr < len(word); chr++ {
					LineStandard = ((int(word[chr]) - 32) * 9) + (1 * Linenum)
					if LineStandard < 0 || LineStandard > 856 {
						fmt.Println("Please write in English")
						return
					}
					Line = Line + Standard[LineStandard]
				}
				fmt.Println(Line)
				Line = ""
			}
		} else if len(word) == 0 {
			NewLine++
			if str == (len(TextArray)-1) && NewLine == len(TextArray){
				return
			}
			fmt.Println()
		}
	}
}

func main() {
	// Read the input
	input := os.Args
	if len(input) > 2 {
		fmt.Println("Please Enter One Argument Only")
	} else if len(input) < 2 {
		fmt.Println("Please Enter An Argument")
	} else {
		Text := string(input[1])
		TextArray := strings.Split(Text, "\\n")

		// Open the file
		file, err := os.Open("standard.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		// Create a scanner for the file
		scanner := bufio.NewScanner(file)

		// Read each line of the file
		var Standard []string
		for scanner.Scan() {
			Standard = append(Standard, scanner.Text())
		}
		// To Print
		PrintString(Standard, TextArray)
	}
}
