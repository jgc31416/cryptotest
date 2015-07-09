/*

Encrypts and decripts vigenere encoded strings

*/
package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("I'm going to crack one time pad")
	byteText := readFile("cipheredTextTimePad.txt")
	fmt.Println(byteText)
	var i byte
	for i = 1; i < 255; i++ {
		for column := 0; column < len(byteText[1]); column++ {
			if checkTransform(byteText, column, i) == true {
				fmt.Printf("\nByte %d makes sense for column %d ", i, column)
			} else {
				//fmt.Printf("Bad %d %d", i, column)
			}
		}
	}
}

func checkTransform(byteText [][]byte, column int, index byte) bool {
	byteColumn := make([]byte, len(byteText))
	for i := 0; i < len(byteText); i++ {
		byteColumn[i] = byteText[i][column] ^ index
	}

	result := hasGoodChars(byteColumn)
	if result == true {
		fmt.Print(string(byteColumn))
	}
	return result
}

func hasGoodChars(text []byte) bool {
	for i := 0; i < len(text); i++ {
		if text[i] < 32 || text[i] > 122 {
			return false
		}
	}

	//Check counters
	return true
}

func readFile(fileName string) [][]byte {
	byteText := make([][]byte, 100)
	var byteLine []byte
	f, err := os.Open(fileName)
	check(err)
	defer f.Close()
	reader := bufio.NewReader(f)
	// Get all the lines converted to byte arrat
	line, err := reader.ReadString('\n')
	i := 0
	for err == nil {
		byteLine, _ = hex.DecodeString(strings.Trim(line, "\n "))
		fmt.Print(byteLine)
		byteText[i] = byteLine
		line, err = reader.ReadString('\n') // 0x0A separator = newline
		i++
	}
	byteLine, _ = hex.DecodeString(line)
	byteText[i] = byteLine

	byteText = byteText[:i+1]

	return byteText
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
