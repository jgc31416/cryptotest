/*

Encrypts and decripts vigenere encoded strings

*/
package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("I'm going to crack one time pad")
	byteText := readFile("cipheredTextTimePad.txt")
	fmt.Println(byteText)
	var i byte
	solution := []byte{242, 26, 4, 155, 208, 115, 35, 200, 57, 152,
		206, 9, 14, 188, 134, 218, 201, 224, 57, 137,
		42, 95, 114, 103, 131, 165, 97, 253, 37, 238,
		48}

	checkSolution(byteText, solution)
	byteText = readFile("cipheredTextTimePad.txt")
	// Look for spaces
	//findSpaceColumns(byteText)

	// Look for remaining
	for i = 1; i < 255; i++ {
		for column := 0; column < len(byteText[1]); column++ {
			if checkTransform(byteText, column, i) == true {
				fmt.Printf(" byte %d makes sense for column %d \n", i, column)
			} else {
				//fmt.Printf("Bad %d %d", i, column)
			}
		}
	}
}

func checkSolution(text [][]byte, solution []byte) {

	for col := 0; col < len(solution); col++ {
		for row := 0; row < 7; row++ {
			text[row][col] = text[row][col] ^ solution[col]
		}
	}
	for row := 0; row < 7; row++ {
		fmt.Println(string(text[row]))
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
	result := true
	//alphaChars := 0
	re := regexp.MustCompile(`[[:lower:]]|[[:blank:]]|,|\.|;|\!|\?`)
	for i := 0; i < len(text); i++ {
		if text[i] < 32 || text[i] > 122 {
			result = false
		}
	}

	if result == true {
		//fmt.Print("\n Found strings")
		//fmt.Print(re.FindAllStringIndex(string(text), -1))
		//fmt.Printf(" - %s\n", string(text))
		if len(re.FindAllStringIndex(string(text), -1)) < 7 {
			result = false
		}

	}
	return result
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
