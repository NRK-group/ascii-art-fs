package banner

import (
	"fmt"
	"strings"
)

/*
	transform the string to art
*/
func AsciiToArt(arg1, ban string) {
	p := AsciiMap(ban)
	for i := 1; i <= 8; i++ {
		for _, value := range arg1 {
			fmt.Print(p[int((value))][i])
		}
		fmt.Println("")
	}
}

/*
	conditions
*/

func PrintAsciiArt(arg1, ban string) {
	switch arg1 {
	case "":
		return
	case "\\n":
		fmt.Println("")
		return
	default:
		argSplit := strings.Split(arg1, "\\n")
		for _, word := range argSplit {
			if word != "" {
				AsciiToArt(word, ban)
			} else {
				fmt.Println("")
			}
		}
	}
}
