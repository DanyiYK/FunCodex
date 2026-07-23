package utils

import (
	"strings"
)

var upperSimilarCharacters = map[string]string{
	"У": "Y",
	"К": "K",
	"Е": "E",
	"Х": "X",
	"А": "A",
	"Р": "P",
	"О": "O",
	"С": "C",

	// Exclusive characters
	"М": "M",
	"Т": "T",
	"Н": "H",
	"В": "B",
	"З": "3",
}

var lowerSimilarCharacters = map[string]string{
	"у": "y",
	"к": "k",
	"е": "e",
	"х": "x",
	"а": "a",
	"р": "p",
	"о": "o",
	"с": "c",
}

var specialSimilarCharacters = map[string]string{
	" ": " ",
}

/*
Test characters are characters that are nearly similiar to other characters
But might be displayed differently with other fonts
*/
var testCharacters = map[string]string{
	"ԁ": "d",
	"һ": "h",
	"і": "i",
	"ј": "j",
	"ѕ": "ѕ",
	"ⅼ": "l", // CHANGE THIS
	"𝗆": "m", // CHANGE THIS
	"𝖿": "f", // CHANGE THIS
	"𝗅": "l", // CHANGE THIS
}

var RegisteredCharmaps = [4]map[string]string{
	upperSimilarCharacters,
	lowerSimilarCharacters,
	specialSimilarCharacters,
	testCharacters,
}

var EncodableCharacters []string = strings.Split("abcdefghijklmnopqrstuvwxyz :3.,", "")

var EmptyCharBit = len(EncodableCharacters) // Special bit that tells the decoder to stop and return the result

const BitsPerCharacter = 5 // Number of bits that makes a char
