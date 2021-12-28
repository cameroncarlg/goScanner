package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"unicode"
)

var counter int
var mySlice []byte

type Token struct {
	tokenType  rune
	tokenValue string
}

var tokenQueue []Token

const LETTER rune = 0
const DIGIT rune = 1
const UNKNOWN rune = 99

const INTEGER rune = 13
const ID_STMT rune = 12
const IN_STMT rune = 10
const LET_STMT rune = 11
const THEN_STMT rune = 21
const IF_STMT rune = 22
const IS_ZERO rune = 23
const MINUS rune = 24
const LEFT_PAREN rune = 25
const RIGHT_PAREN rune = 26
const EQU_OP rune = 27
const COMMA rune = 28
const ELSE_STMT rune = 29

var charClass rune
var lexeme string

var nextChar byte
var token int
var nextToken rune
var turnOff bool

func Scanner() {

	var filename string
	fmt.Print("Enter file: ")
	fmt.Scanln(&filename)
	fmt.Println()

	data, err := ioutil.ReadFile(filename)
	//data, err := ioutil.ReadFile("/home/cameron/let_lang_proj/front2.in")
	if err == io.EOF {
		fmt.Println("Reading file finished...")
		return
	}

	fmt.Println(string(data))

	tokenQueue = []Token{}

	mySlice = data[:]
	my_getChar()
	for !turnOff {
		lex()
	}

	for _, v := range tokenQueue {
		fmt.Println("tok: ", v)
	}
	fmt.Println()

}

func my_addChar() {
	tmpString := string(nextChar)
	lexeme += tmpString
}

func my_getChar() {
	nextChar = mySlice[counter]

	if unicode.IsLetter(rune(nextChar)) {
		charClass = LETTER
	} else if unicode.IsDigit(rune(nextChar)) {
		charClass = DIGIT
	} else {
		charClass = UNKNOWN
	}

	// Increment global counter
	counter++
}

func getNonBlank() {
	for unicode.IsSpace(rune(nextChar)) {
		my_getChar()
	}
}

func lex() {
	lexeme = " "
	getNonBlank()
	switch charClass {
	case LETTER:
		my_addChar()
		my_getChar()
		for charClass == LETTER || charClass == DIGIT {
			my_addChar()
			my_getChar()
		}
		nextToken = ID_STMT
	case DIGIT:
		my_addChar()
		my_getChar()
		for charClass == DIGIT {
			my_addChar()
			my_getChar()
		}
		nextToken = INTEGER
	case UNKNOWN:
		Lookup(nextChar)
		// LOOK UP: ( ) , =,
		my_getChar()

	}

	switch lexeme {
	case " let":
		nextToken = LET_STMT
	case " in":
		nextToken = IN_STMT
	case " iszero":
		nextToken = IS_ZERO
	case " if":
		nextToken = IF_STMT
	case " then":
		nextToken = THEN_STMT
	case " else":
		nextToken = ELSE_STMT
	case " minus":
		nextToken = MINUS
	}

	// fmt.Println("Next token is:", nextToken, ", Next lexeme is:", lexeme)

	tokenQueue = append(tokenQueue, Token{nextToken, lexeme[1:]})

	if counter == len(mySlice) {
		turnOff = true
	}

}

func Lookup(nextChar byte) {
	switch nextChar {
	case '(':
		my_addChar()
		nextToken = LEFT_PAREN
	case ')':
		my_addChar()
		nextToken = RIGHT_PAREN
	case ',':
		my_addChar()
		nextToken = COMMA
	case '=':
		my_addChar()
		nextToken = EQU_OP
	}
}
