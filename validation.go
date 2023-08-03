package ascii

import (
	"fmt"
	"os"
	"strings"
)

func Validation() string {
	//To check the number of arguments
	val := "yes"
	a := 3
	b := 1
	Test := -1
	flag2 := false
	if len(os.Args) <= 4 && len(os.Args) > 1 {
		Test = strings.Index(os.Args[1], "--output=")
		if Test == 0 {
			val = "output"
			if len(os.Args) != 2 && len(os.Args[1]) > 9 && strings.Index(os.Args[1], ".txt") != -1 && len(os.Args[1]) == strings.Index(os.Args[1], ".txt")+4 {
				a++
				b++
			} else {
				Error()
			}

		} else {
			Test = strings.Index(os.Args[1], "--color=")
		}
		FlgForcolor := false
		if strings.Index(os.Args[1], "--color=") == 0 {
			val = "color"
			if  len(os.Args) == 4 && len(os.Args[3])> len(os.Args[2]) && os.Args[3] != "standard" && os.Args[3] != "shadow" && os.Args[3] != "thinkertoy" {
				d := strings.Index(os.Args[3], os.Args[2])
				if d == -1 {
					Error()
				} else {
					FlgForcolor = true
					val = "colorWletter"
				}
			}
			color := strings.ToLower(strings.TrimPrefix(os.Args[1], "--color="))
			if CheckColor(color) != "NO" {
				flag2 = true
			}
			if len(os.Args) != 2 && flag2 == true && len(os.Args[1]) > 8 {
				a++
				b++
			} else {
				Error()
			}
		}
		//To validate the character of the strings in ascii range only
		for g := 0; g < len(os.Args[b]); g++ {
			if os.Args[b][g] > 126 || os.Args[b][g] < 32 {
				fmt.Println("ERROR: ascii letters only")
				os.Exit(0)
			}
		}
		if FlgForcolor {
			for g := 0; g < len(os.Args[3]); g++ {
				if os.Args[3][g] > 126 || os.Args[3][g] < 32 {
					fmt.Println("ERROR: ascii letters only")
					os.Exit(0)
				}
			}
		}

		if len(os.Args) == a && !FlgForcolor  {
			fmt.Println("Hi1")
			FontType := strings.ToLower(os.Args[a-1])
			if FontType != "standard" && FontType != "shadow" && FontType != "thinkertoy" {
				Error()
			}
		} else if len(os.Args) == 2 || (Test == 0 && len(os.Args) == 3) || FlgForcolor {
			// fmt.Println("Hi2")
		} else {
			fmt.Println("Hi3")
			Error()
		}
	} else {
		Error()
	}
	if len(os.Args[b]) == 0 {
		return "no"
	}
	return val
}

func Error() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\"")
	os.Exit(0)
}
