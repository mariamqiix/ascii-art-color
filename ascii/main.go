package main

import (
	"ascii"
	"fmt"
	"os"
	"strings"
)

func main() {
	validation := ascii.Validation()
	a := 1
	b := 2
	if validation != "no" {
		if validation == "output" || validation == "color" {
			a++
			b++
		} else if validation == "colorWletter" || validation == "colorWletterWfont" {
			a += 2
			b += 2
		}
		fmt.Println(validation)
		WordsInArr := strings.Split(os.Args[a], "\\n")
		fileName := "standard"
		if len(os.Args) == 3 && validation != "output" && validation != "color" {
			fileName = strings.ToLower(os.Args[b])
		} else if len(os.Args) == 4 && validation != "colorWletter" {
			fileName = strings.ToLower(os.Args[b])
		} else if validation == "colorWletterWfont" {
			fileName = strings.ToLower(os.Args[4])
		}
		if ascii.OnlyContains(os.Args[a], "\\n") {
			WordsInArr = WordsInArr[:len(WordsInArr)-1]
		}
		FirstWord := true
		for l := 0; l < len(WordsInArr); l++ {
			var Words [][]string
			Text1 := WordsInArr[l]
			Text1 = strings.ReplaceAll(Text1, "\\t", "   ")
			for j := 0; j < len(Text1); j++ {
				Words = append(Words, ascii.ReadLetter(Text1[j], fileName))
			}
			if validation == "output" {
				ascii.WriteFile(Words, FirstWord)
				FirstWord = false
			} else if validation == "color" || validation == "colorWletter" || validation == "colorWletterWfont" {
				letter := Text1
				letterLen := "NO"
				color := ascii.CheckColor(strings.ToLower(strings.TrimPrefix(os.Args[1], "--color=")))
				if validation == "colorWletter" || validation == "colorWletterWfont" {
					letterLen = os.Args[a-1]
				}
				PrintWithColor(Words, color, letter, letterLen)
			} else {
				for w := 0; w < 8; w++ {
					if len(Text1) == 0 {
						break
					}
					for n := 0; n < len(Words); n++ {
						fmt.Print(Words[n][w])
					}
					if w+1 != 8 {
						fmt.Println()
					}
				}
				fmt.Println()
			}
		}
	}
}

func PrintWithColor(Words [][]string, color, letterA, letterLenA string) {
	TheWord := letterA
	letterLen := len(letterLenA)
	colorB := color
	var nbr []int
	wordA := ""
	nn := len(letterA)
	c := 0
	for i := 0; i < len(letterA); i++ {
		before, after, found := strings.Cut(TheWord, letterLenA)
		if found == true {
			if len(after) != 0 {
				wordA = after
				c = len(after)
			} else if len(before) != 0 {
				wordA = before
				c = len(after)
			}
			nn = len(letterA) - c - len(letterLenA)
			nbr = append(nbr, nn)
		} else {
			break
		}
		TheWord = wordA
	}
	fmt.Println(nbr)
	FlagB := false
	for w := 0; w < 8; w++ {
		d := 0
		//strings.Index(letterA, letterLenA)
		// if len(Text1) == 0 {
		// 	break
		// }
		a := 0
		for n := 0; n < len(Words); n++ {
			if letterLenA != "NO" {
				if len(nbr) != 0 {
					chngeLetter := nbr[d]
					if chngeLetter == n {
						FlagB = true
					}
					if FlagB {
						colorB = color
						a++
						if a == letterLen {
							FlagB = false
							a = 0
							if d+1 < len(nbr) {
								d++
							}
						}
					} else {
						colorB = "\033[0m"
					}
				}
			}
			fmt.Print(colorB, Words[n][w])
		}
		if w+1 != 8 {
			fmt.Println()
		}
	}
	fmt.Println()
}
