package utils

import (
	"fmt"
	"math"
	"strings"
)

func ArrayFind(arrayToIterate []string, valueToFind string) int {
	for i, v := range arrayToIterate {
		if v == valueToFind {
			return i
		}
	}

	return -1
}

func TableFind(mapToIterate map[string]string, valueToFind string) string {
	for k, v := range mapToIterate {
		if k == valueToFind {
			return v
		}
	}

	return ""
}

func IntegerToBinary(number int, size int) [BitsPerCharacter]int {
	//fmt.Print(number)
	binaryNumber := fmt.Sprintf("%b", number)
	splittedBinaryNumber := strings.Split(binaryNumber, "")

	returnVal := [BitsPerCharacter]int{}

	// Add zeros to it to match the requested size
	if len(binaryNumber) < BitsPerCharacter {
		for i := 0; i < (BitsPerCharacter - len(splittedBinaryNumber)); i++ {
			binaryNumber = "0" + binaryNumber
		}
	}

	fmt.Println(binaryNumber)

	for i, bit := range strings.Split(binaryNumber, "") {
		if bit == "0" {
			returnVal[i] = 0
			continue
		}

		returnVal[i] = 1
	}

	return returnVal
}

func BinaryToInteger(binaryNumber [BitsPerCharacter]int) int {
	returnVal := 0

	for i := BitsPerCharacter - 1; i >= 0; i-- {
		if binaryNumber[i] == 0 {
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

func PopHiddenTextBinary(intArray [][BitsPerCharacter]int) (array [][BitsPerCharacter]int, value [BitsPerCharacter]int) {
	newArray := [][BitsPerCharacter]int{}
	lastValue := intArray[len(intArray)-1]

	for i := 0; i < len(intArray)-1; i++ {
		newArray[i] = intArray[i]
	}

	return newArray, lastValue
}

/*func PopParsing(intArray [BitsPerCharacter]int) (array []int, value int) {
	newArray := [BitsPerCharacter]int{}
	lastValue := intArray[len(intArray)-1]

	for i := 0; i < len(intArray)-1; i++ {
		newArray[i] = intArray[i]
	}

	return newArray, lastValue
}*/
