package asciiart

import (
	"asciiartfs/banner"
	"fmt"
	"os"
)

func MainFormatError() {
	fmt.Println("Usage: go run . [STRING]")
	fmt.Println("")
	fmt.Println("EX: go run . something")
}

func AsciiArtMain() {
	if len(os.Args[1:]) == 1 {
		banner.PrintAsciiArt(os.Args[1], banner.Standard)
	} else {
		MainFormatError()
	}
}
