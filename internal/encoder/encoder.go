package encoder

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Danyiyk/FunCodex/internal/utils"
)

func getEncodedLetter(charMap map[string]string, letter string, bit int) string {
	for k, v := range charMap {
		if k == letter || v == letter {
			if bit == 0 {
				return k
			} else {
				return v
			}
		}
	}

	return ""
}

func Encode(str string, hiddenText string) string {
	splittedMainString := strings.Split(str, "")

	splittedString := strings.Split(hiddenText, "")
	returnValue := ""
	hiddenTextBinary := [][utils.BitsPerCharacter]int{}

	// Check if length match
	_, maxCharacters := utils.GetAvailableSpace(str)
	hiddenTextLength := len(splittedString)

	for i := range maxCharacters {
		if i >= hiddenTextLength {
			hiddenTextBinary = append(hiddenTextBinary, utils.IntegerToBinary(utils.EmptyCharBit, utils.BitsPerCharacter))
			break
		}

		letter := splittedString[i]
		found := utils.ArrayFind(utils.EncodableCharacters, letter)

		if found != -1 {
			hiddenTextBinary = append(hiddenTextBinary, utils.IntegerToBinary(found, utils.BitsPerCharacter))
		} else {
			fmt.Printf("[Warning] Invalid character in hidden string: %s\n", letter)
		}
	}

	slices.Reverse(hiddenTextBinary)

	if hiddenTextLength > maxCharacters {
		fmt.Printf("[Warning] Hidden text is too long to be hidden in this string, it will be cut!")
	}

	var parsing [utils.BitsPerCharacter]int
	var parsingPosition int
	cursor := len(hiddenTextBinary)

	for i := range splittedMainString {
		letter := splittedMainString[i]

		bitValue, letterCharmap := utils.GetBit(letter)

		if cursor != 0 && parsingPosition == 0 {
			cursor--
			parsing = hiddenTextBinary[cursor]
			parsingPosition = utils.BitsPerCharacter

			fmt.Println(parsing)
		}

		if bitValue == -1 || cursor == 0 {
			returnValue = returnValue + letter
			continue
		}

		parsingPosition--
		returnValue = returnValue + getEncodedLetter(letterCharmap, letter, parsing[parsingPosition])
	}

	return returnValue
}
