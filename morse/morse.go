package morse

import (
	"fmt"
	"strings"
	"regexp"
)

func Decode(s string, alphabet map[string]string, letterSeparator string, wordSeparator string) (string, error) {

	res := ""
	for _, part := range strings.Split(s, letterSeparator) {
		found := false
		for key, val := range alphabet {
			if val == part {
				res += key
				found = true
				break
			}
		}
		if part == wordSeparator {
			res += " "
			found = true
		}
		if found == false {
			return res, fmt.Errorf("unknown character " + part)
		}
	}
	return res, nil
}

func Encode(s string, alphabet map[string]string, letterSeparator string, wordSeparator string) string {

	res := ""
	s = strings.ToLower(s)
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	s = reg.ReplaceAllString(s, "")
	for _, part := range s {
		p := string(part)
		if p == " " {
			if wordSeparator != "" {
				res += wordSeparator + letterSeparator
			}
		} else if morseITU[p] != "" {
			res += morseITU[p] + letterSeparator
		}
	}
	return strings.TrimSpace(res)
}

// DecodeITU translates international morse code (ITU) to text
func DecodeITU(s string) (string, error) {
	return Decode(s, morseITU, " ", "/")
}

// EncodeITU translates text to international morse code (ITU)
func EncodeITU(s string) string {
	return Encode(s, morseITU, " ", "/")
}

// LooksLikeMorse returns true if string seems to be a morse encoded string
func LooksLikeMorse(s string) bool {

	if len(s) < 1 {
		return false
	}
	for _, b := range s {
		if b != '-' && b != '.' && b != ' ' {
			return false
		}
	}
	return true
}

var (
	morseITU = map[string]string{
		"a":  ".-",
		"b":  "-...",
		"c":  "-.-.",
		"d":  "-..",
		"e":  ".",
		"f":  "..-.",
		"g":  "--.",
		"h":  "....",
		"i":  "..",
		"j":  ".---",
		"k":  "-.-",
		"l":  ".-..",
		"m":  "--",
		"n":  "-.",
		"o":  "---",
		"p":  ".--.",
		"q":  "--.-",
		"r":  ".-.",
		"s":  "...",
		"t":  "-",
		"u":  "..-",
		"v":  "...-",
		"w":  ".--",
		"x":  "-..-",
		"y":  "-.--",
		"z":  "--..",
		"ä":  ".-.-",
		"ö":  "---.",
		"ü":  "..--",
		"Ch": "----",
		"0":  "-----",
		"1":  ".----",
		"2":  "..---",
		"3":  "...--",
		"4":  "....-",
		"5":  ".....",
		"6":  "-....",
		"7":  "--...",
		"8":  "---..",
		"9":  "----.",
		".":  ".-.-.-",
		",":  "--..--",
		"?":  "..--..",
		"!":  "..--.",
		":":  "---...",
		"\"": ".-..-.",
		"'":  ".----.",
		"=":  "-...-",
	}
)
