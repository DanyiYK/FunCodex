package utils

import (
	"fmt"
	"math"
	"strings"
)

func TableFind(mapToIterate map[string]string, valueToFind string) string {
	for k, v := range mapToIterate {
		if k == valueToFind {
			return v
		}
	}

	return ""
}

func IntegerToBinary(number int, size int) string {
	binaryNumber := fmt.Sprintf("%b", number)

	// Add zeros to it to match the requested size
	if len(binaryNumber) < size {
		for i := 0; i < (size - len(binaryNumber)); i++ {
			binaryNumber = "0" + binaryNumber
		}
	}

	return binaryNumber
}

func BinaryToInteger(binaryNumber string) int {
	splittedString := strings.Split(binaryNumber, "")
	returnVal := 0

	for i := len(splittedString) - 1; i >= 0; i-- {
		if splittedString[i] == "0" {
			continue
		}

		returnVal += int(math.Pow(2, float64(i)))
	}

	return returnVal
}

/*
Returns 0 if it's a key of the map
Returns 1 if it's a value of the map
Returns -1 if it's neither
*/
func IsKeyOrValue(mapToCheck map[string]string, value string) int {
	for k, v := range mapToCheck {
		if k == value {
			return 0
		} else if v == value {
			return 1
		}
	}

	return -1
}

func GetBit(char string) (int, map[string]string) {
	for _, charMap := range RegisteredCharmaps {
		bitValue := IsKeyOrValue(charMap, char)

		if bitValue != -1 {
			return bitValue, charMap
		}
	}

	return -1, map[string]string{}
}

func GetAvailableSpace(text string) (encodableCharCount int, hiddenTextSize int) {
	splittedString := strings.Split(text, "")
	count := 0

	for i := 0; i < len(splittedString); i++ {
		letter := splittedString[i]

		bit, _ := GetBit(letter)

		if bit != -1 {
			count++
		}
	}

	return count, int(math.Floor(float64(count) / float64(BitsPerCharacter)))
}
