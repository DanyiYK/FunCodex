package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Danyiyk/FunCodex/internal/decoder"
	"github.com/Danyiyk/FunCodex/internal/encoder"
	"github.com/Danyiyk/FunCodex/internal/utils"
)

func PrintHelp() {
	fmt.Println("Registered charmaps", len(utils.RegisteredCharmaps))
	fmt.Println("Available commands:\n -l(ength) [text]\n -c(rypt) [text] [hidden_text]\n -d(ecrypt) [crypted_text]\n -lc(-lengthcrypt) [text_to_crypt]")
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
		fmt.Println(`"` + decoder.Decode(text) + `"`)

	case "-l":
		encodableCharacters, hiddenTextMaxLength := utils.GetAvailableSpace(text)

		fmt.Printf("Analysis:\n Text: \"%s\"\n Encodable characters: %d\n Hidden text max length: %d characters\n", text, encodableCharacters, hiddenTextMaxLength)

	case "-lc":
		encodableCharacters, hiddenTextMaxLength := utils.GetAvailableSpace(text)

		fmt.Printf("There are %d encodable characters in this sentence, hidden text can be up to %d characters, would you like to continue? ", encodableCharacters, hiddenTextMaxLength)

		if hiddenTextMaxLength == 0 {
			fmt.Println("\nThe string is not long enough to store any character")
			break
		}

		fmt.Printf("[Y/n]: ")

		var response string

		fmt.Scanf("%s", &response)

		splitted := strings.Split(response, "")
		if len(splitted) != 0 && splitted[0] == "n" {
			break
		}

		var hidden_text string

		fmt.Printf("Insert hidden text (max of %d characters):\n", hiddenTextMaxLength)

		for i := 0; i < hiddenTextMaxLength; i++ {
			fmt.Printf("▮")
		}

		fmt.Printf("\r")

		scanner := bufio.NewScanner(os.Stdin)

		if !scanner.Scan() {
			break
		}

		hidden_text = scanner.Text()

		fmt.Printf("Encoding the %d character long hidden text\n", len(strings.Split(hidden_text, "")))

		fmt.Println(`Encoded string: "` + encoder.Encode(text, hidden_text) + `"`)
	default:
		PrintHelp()
	}
}
