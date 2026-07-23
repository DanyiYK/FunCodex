package main

import (
	"fmt"
	"os"

	"github.com/Danyiyk/FunCodex/internal/decoder"
	"github.com/Danyiyk/FunCodex/internal/encoder"
	"github.com/Danyiyk/FunCodex/internal/utils"
)

func PrintHelp() {
	fmt.Println("Registered charmaps", len(utils.RegisteredCharmaps))
	fmt.Println("Available commands:\n -l(ength) [text]\n -c(rypt) [text] [hidden_text]\n -d(ecrypt) [crypted_text]")
}

func main() {
	var mode string
	var text string

	if len(os.Args) < 3 {
		fmt.Println("Mismatching argument count!")
		PrintHelp()

		return
	}

	mode = os.Args[1]

	if len(os.Args) >= 3 {
		text = os.Args[2]
	} else {
		text = ""
	}

	switch mode {
	case "-c":
		if len(os.Args) < 4 {
			fmt.Println("-c [text] [text to encrypt]")
			return
		}

		hidden_text := os.Args[3]

		fmt.Println(encoder.Encode(text, hidden_text))

	case "-d":
		fmt.Println(decoder.Decode(text))

	case "-l":
		encodableCharacters, hiddenTextMaxLength := utils.GetAvailableSpace(text)

		fmt.Printf("Analysis:\n Text: \"%s\"\n Encodable characters: %d\n Hidden text max length: %d characters\n", text, encodableCharacters, hiddenTextMaxLength)
	default:
		PrintHelp()
	}
}
