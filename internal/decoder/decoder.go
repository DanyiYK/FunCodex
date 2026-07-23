package decoder

import (
	"strings"

	"github.com/Danyiyk/FunCodex/internal/utils"
)

func Decode(encodedText string) string {
	returnVal := ""
	foundBinary := [][utils.BitsPerCharacter]int{}

	splittedText := strings.Split(encodedText, "")

	filling := [utils.BitsPerCharacter]int{}
	cursor := 0

	for _, letter := range splittedText {
		if cursor == 5 {
			cursor = 0
			foundBinary = append(foundBinary, filling)

			filling = [utils.BitsPerCharacter]int{}
		}

		found, _ := utils.GetBit(letter)

		if found != -1 {
			filling[cursor] = found
			cursor++
		}

	}

	for _, binaryToDecode := range foundBinary {
		value := utils.BinaryToInteger(binaryToDecode)

		if value == utils.EmptyCharBit {
			break
		}

		returnVal += utils.EncodableCharacters[value]
	}

	return returnVal
}
