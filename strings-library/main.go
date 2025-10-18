package main

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	// var stringVar1 string = "Learning Go is Fun and Exciting\n"

	//Basic string Operation
	stringVar1 := "Learning Go is Fun and Exciting\n"
	stringVar2 := "Learning"

	// string comparison
	fmt.Println(strings.Compare(stringVar2, stringVar1))

	fmt.Println(strings.Contains(stringVar1, "Exciting"))
	fmt.Println(strings.ContainsAny(stringVar1, "F65432"))

	fmt.Println(strings.Count(stringVar1, "i"))
	bf, af, fo := strings.Cut(stringVar1, "is")

	fmt.Printf("The before is: %s\nand the after is: %s\nand the found is: %t\n", bf, af, fo)

	fmt.Println(strings.EqualFold("GO", "go"))

	fmt.Println(strings.Fields(stringVar1))
	stringVar3 := strings.Fields(stringVar1)

	// for i := 0; i < len(stringVar3); i++ {
	// 	fmt.Printf("The value of the : %v\nand the corresponding value is :%s\n", i, stringVar3[i])
	// }

	//another way to iterate over the string slice

	for i, value := range stringVar3 {
		fmt.Printf("The value at index : %v\nand it's corresponding value is : %s\n", i, value)
	}

	function := func(character rune) bool {

		return !unicode.IsLetter(character) && !unicode.IsNumber(character)
	}

	fmt.Println(strings.FieldsFunc("Strg1;::;;;;;;(*&^%$#)   String2, String3 ", function))

	fmt.Println(strings.HasSuffix(stringVar2, "g"))
	fmt.Println(strings.HasPrefix(stringVar1, "L"))

	//using the join function in the string library
	mySlice := []string{"Here", "Iam", "Learning", "Go", "Language"}
	fmt.Printf("The type of the mySlice var is : %T  before Join\n ", mySlice)
	fmt.Println(strings.Join(mySlice, " <--> "))

	joinedSlice := strings.Join(mySlice, " <--> ")
	fmt.Printf("The type of the joinedSlice var is : %T  after Join function\n ", joinedSlice)

	fmt.Println("\n" + strings.Repeat(stringVar1, 3))

	var indexString string = "Learning Go is Fun and Exciting Go is Fun"
	fmt.Println(strings.Index(indexString, "Fun"))
	fmt.Println(strings.LastIndex(indexString, "Fun"))

	//find the index of any of the characters in the given string at first occurrence and returns the position of the  index of the first occurrence
	fmt.Println(strings.IndexAny(indexString, "is"))

	//find the index of any of the characters in the given string at last occurrence and returns the position of the  index of the last occurrence
	fmt.Println(strings.LastIndexAny(indexString, "is"))

	//find the index of the given byte character in the string and returns the position of the index of the first occurrence Note:a byte is an alias for uint8 and Note: that a byte can only represent ASCII characters and Note: Unicode characters cannot be represented by a single byte and note: for Unicode characters use rune type instead and note:rune is an alias for int32 and note: rune can represent all Unicode characters and note: Unicode characters may require more than one byte to be represented.

	// note: 'i' is an ASCII character and can be represented by a single byte using the single quotes only and not double quotes and double quotes are used to represent string literals
	fmt.Println(strings.IndexByte(indexString, 'i'))

	//find the index of the given byte character in the string and returns the position of the index of the last occurrence
	fmt.Println(strings.LastIndexByte(indexString, 'i'))

	isChineseCharacter := func(runeVar rune) bool {

		return unicode.Is(unicode.Han, runeVar)
	}

	fmt.Println(strings.IndexFunc("Hello, 你好", isChineseCharacter))
	fmt.Println(strings.LastIndexFunc("Hello, 你好", isChineseCharacter))

	fmt.Println(strings.IndexRune(indexString, 'F'))

	fmt.Println(strings.ContainsRune(indexString, 'F'))

	fmt.Println(strings.Replace(indexString, "Fun", "Nice", -1000000000000000000))
	fmt.Println(strings.ReplaceAll(indexString, "Go", "Golang"))

	splitString := "A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z"
	fmt.Println(strings.Split(splitString, ","))
	fmt.Println(strings.SplitN(splitString, ",", 3))
	fmt.Println(strings.SplitN(splitString, ",", -1))

	fmt.Println(strings.SplitAfter(splitString, ","))
	fmt.Println(strings.SplitAfterN(splitString, ",", 4))

	englishCaser := cases.Title(language.English)
	titledString := englishCaser.String(stringVar1)
	fmt.Println(titledString)

	//to title case
	totitledString := strings.ToTitle(stringVar1)
	fmt.Println(totitledString)

	//to lower case
	lowerString := strings.ToLower(stringVar1)
	fmt.Println(lowerString)

	//to upper case
	upperString := strings.ToUpper(stringVar1)
	fmt.Println(upperString)

	specialeString := "São Paulo is a large city. Here are more characters: à, é, í, õ, ç, ñ, ü"

	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, specialeString))

	loweSpecialString := strings.ToLowerSpecial(unicode.TurkishCase, "SÃO PAULO İS A LARGE CİTY. HERE ARE MORE CHARACTERS: À, É, Í, Õ, Ç, Ñ, Ü")

	fmt.Println(loweSpecialString)

	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "são paulo is a large city. here are more characters: à, é, í, õ, ç, ñ, ü"))

}
