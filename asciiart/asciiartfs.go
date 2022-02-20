package asciiart

import (
	"asciiartfs/banner"
	"fmt"
	"os"
)

func FSFormatError() {
	fmt.Println("Usage: go run . [STRING] [BANNER]")
	fmt.Println("")
	fmt.Println("EX: go run . something standard")
}

func AsciiArtFS() {
	if len(os.Args[1:]) == 2 {
		switch os.Args[2] {
		case "standard":
			banner.PrintAsciiArt(os.Args[1], banner.Standard)
		case "shadow":
			banner.PrintAsciiArt(os.Args[1], banner.Shadow)
		case "thinkertoy":
			banner.PrintAsciiArt(os.Args[1], banner.Thinkertoy)
		default:
			FSFormatError()
		}
	} else {
		FSFormatError()
	}
}
