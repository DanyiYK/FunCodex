package decoder

import (
	"fmt"
	"strings"

	"github.com/Danyiyk/FunCodex/internal/utils"
)

func Decode(encodedText string) string {
	var returnVal strings.Builder
	foundBinary := [][utils.BitsPerCharacter]int{}

	splittedText := strings.Split(encodedText, "")

	filling := [utils.BitsPerCharacter]int{}
	cursor := 0

	for i, letter := range splittedText {
		found, _ := utils.GetBit(letter)

		if found != -1 {
			filling[cursor] = found
			cursor++
		}

		if cursor != 5 {
			continue
		}

		cursor = 0
		foundBinary = append(foundBinary, filling)

		if i != len(splittedText)-1 {
			filling = [utils.BitsPerCharacter]int{}
		}
	}

	for i, binaryToDecode := range foundBinary {
		value := utils.BinaryToInteger(binaryToDecode)

		if value == utils.EmptyCharBit {
			fmt.Println("break signal received", i, value, utils.EmptyCharBit)
			break
		}

		returnVal.WriteString(utils.EncodableCharacters[value])
	}

	return returnVal.String()
}
