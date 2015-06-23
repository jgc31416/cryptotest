/*

Encrypts and decripts vigenere encoded strings

*/
package main

import (
	"encoding/hex"
	"fmt"
	"math"
)

func main() {
	var textCiphered string
	var byteText []byte
	var keyLenght int
	fmt.Println("I'm going to crack vigenere")

	fmt.Scanln(&textCiphered)
	byteText, _ = hex.DecodeString(textCiphered)

	keyLenght = getKeyLenght(byteText)
	fmt.Printf("\nKey length: %d\n", keyLenght)
	guessKey(keyLenght, byteText)

	//Key found after fiddling around a bit
	key := []byte{
		186, 31, 145, 178, 83, 205, 62,
	}
	fmt.Print(string(decrypt(byteText, key)))
}

/*
func main() {
	//var textCiphered string
	var byteText []byte
	var keyLenght int
	fmt.Println("I'm going to apply and crack vigenere")

	var plainText, key string
	fmt.Scanln(&plainText)
	key = "Maqsu9g" // no use
	byteText = encrypt([]byte(plainText), []byte(key))

	fmt.Scanln(&textCiphered)
	byteText, _ = hex.DecodeString(textCiphered)

	keyLenght = getKeyLenght(byteText)
	fmt.Printf("\nKey length: %d\n", keyLenght)
	guessKey(keyLenght, byteText)

}
*/

/*
Encrypts a given text
*/
func encrypt(plainText []byte, key []byte) (cipher []byte) {
	fmt.Print("Ciphering!")
	cipher = make([]byte, len(plainText))
	keyLenght := len(key)
	for i := 0; i < len(plainText); i++ {
		cipher[i] = plainText[i] ^ key[i%keyLenght]
	}
	return cipher
}

func decrypt(text []byte, key []byte) (plainText []byte) {
	fmt.Print("DeCiphering!")
	plainText = make([]byte, len(text))
	keyLenght := len(key)
	for i := 0; i < len(text); i++ {
		plainText[i] = text[i] ^ key[i%keyLenght]
	}
	return plainText
}

func guessKey(keyLength int, text []byte) (plainText string) {
	possibleKey := make([]byte, keyLength)
	var xorText byte
	//Apply xor with 0 to 255 on each key "chunk"
	for i := 0; i < 256; i++ {
		var charsToCheck = make([][]byte, keyLength)
		for n := 1; n < keyLength; n++ {
			charsToCheck[n] = make([]byte, 1)
			charsToCheck[n] = charsToCheck[n][1:]
		}

		//Get all the chars by key positon and xor with char
		for j := 0; j < len(text); j++ {
			xorText = text[j] ^ byte(i)
			charsToCheck[j%keyLength] = append(charsToCheck[j%keyLength], xorText)
		}

		for k := 0; k < keyLength; k++ {
			//Guess if chars found pass conditions
			if hasValidChars(charsToCheck[k]) == true {
				fmt.Printf("\nOffset %d char %s of key position %d passes valid chars ", i, string(i), k)
				fmt.Print(string(charsToCheck[k]))
				possibleKey[k] = byte(i)
			}
		}
	}

	fmt.Printf("\nPossible key: %s \n", string(possibleKey))
	return plainText
}

func hasValidChars(text []byte) (validChars bool) {
	//Sensible frequencies for chars

	freqCharsTable := map[byte]float64{
		//97:  0.085,
		//101: 0.125,
		116: 0.090,
	}

	freqCharsTableCounter := map[byte]float64{}
	// Valid range check
	for i := 0; i < len(text); i++ {
		if text[i] < 32 || text[i] > 127 {
			return false
		} else if text[i] < 58 && text[i] > 47 {
			return false
		}
		freqCharsTableCounter[text[i]]++
	}

	// Valid frequencies check not very reliable
	for key, value := range freqCharsTable {
		txtFreq := freqCharsTableCounter[key] / float64(len(text))
		//fmt.Printf("\nTesting char %d %s freq %f vs value %f", key, string(key), txtFreq, value)
		if txtFreq < value-0.05 || txtFreq > value+0.05 {
			return false
		}
	}
	// This needs a bit of fiddling around
	return true
}

func getKeyLenght(text []byte) (keyLength int) {
	var maxKeyLenght = 13
	var maxDeviation float64
	for i := 3; i <= maxKeyLenght; i++ {
		deviation := getTextDeviation(text, i)
		if deviation > maxDeviation {
			maxDeviation = deviation
			keyLength = i
		}
	}
	return keyLength
}

func getTextDeviation(text []byte, keyLenght int) (deviation float64) {
	freqCharsCounter := make(map[byte]int)
	//Calculate frequencies
	totalChars := 0
	for i := 0; i < len(text); i++ {
		if i%keyLenght == 0 {
			freqCharsCounter[text[i]]++
			totalChars++
		}
	}

	for _, counter := range freqCharsCounter {
		deviation = deviation + math.Pow(float64(counter)/float64(totalChars), 2)
	}

	fmt.Printf("\nFound %d different chars with key lenght %d deviation: %f",
		len(freqCharsCounter),
		keyLenght,
		deviation)

	return deviation
}
