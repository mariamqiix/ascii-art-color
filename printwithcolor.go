package ascii

import (
	"fmt"
	"os"
)

func PrintWithColor(Words [][]string, color, Text1, letterLenA, validation string) {
	letterLen := len(letterLenA)
	colorB := color
	nbr, colors := CheckPostion(Text1, colorB, letterLenA, validation)
	FlagB := false
	for w := 0; w < 8; w++ {
		d := 0
		if len(Text1) == 0 {
			break
		}
		a := 0
		for n := 0; n < len(Words); n++ {
			if letterLenA != "NO!!-" {
				if len(nbr) != 0 {
					chngeLetter := nbr[d]
					if chngeLetter == n {
						FlagB = true
					}
					if FlagB {
						colorB = colors[d]
						a++
						if (a == letterLen && colorB == color) || (validation == "colorW2letter" && a == len(os.Args[4]) && colorB != color) {
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
