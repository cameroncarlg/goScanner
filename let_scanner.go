package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"unicode"
)

type person struct {
	name string
	age  int
}

var counter int
var mySlice []byte

const LETTER int = 0
const DIGIT int = 1
const UNKNOWN int = 99

const INT_LIT int = 10
const IDENT int = 11
const ADD_OP int = 21
const SUB_OP int = 22
const MULT_OP int = 23
const DIV_OP int = 24
const LEFT_PAREN int = 25
const RIGHT_PAREN int = 26

//const EOF int = -1

var charClass int
var lexeme string

var nextChar byte
var token int
var nextToken int
var turnOff bool

func Scanner() {
	data, err := ioutil.ReadFile("/home/cameron/Downloads/front.in")
	if err == io.EOF {
		fmt.Println("Reading file finished...")
		return
	}

	fmt.Println(person{"Bob", 20})

	mySlice = data[:]
	My_getChar()
	for !turnOff {
		Lex()
	}

}

func newPerson(name string) *person {
	p := person{name: name}
	//p.age = 42
	return &p
}

func My_addChar() {
	tmpString := string(nextChar)
	lexeme += tmpString
}

func My_getChar() {
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

func GetNonBlank() {
	for unicode.IsSpace(rune(nextChar)) {
		My_getChar()
	}
}

func Lex() {
	lexeme = " "
	GetNonBlank()
	switch charClass {
	case LETTER:
		My_addChar()
		My_getChar()
		for charClass == LETTER || charClass == DIGIT {
			My_addChar()
			My_getChar()
		}
		nextToken = IDENT
	case DIGIT:
		My_addChar()
		My_getChar()
		for charClass == DIGIT {
			My_addChar()
			My_getChar()
		}
		nextToken = INT_LIT
	case UNKNOWN:
		Lookup(nextChar)
		My_getChar()

	}

	fmt.Println("Next token is: ", nextToken, ", Next lexeme is: ", string(lexeme))

	if counter == len(mySlice) {
		turnOff = true
		fmt.Println("Next token is:  -1 , Next lexeme is:   EOF")
	}

}

func Lookup(nextChar byte) {
	switch nextChar {
	case '(':
		My_addChar()
		nextToken = LEFT_PAREN
	case ')':
		My_addChar()
		nextToken = RIGHT_PAREN
	case '+':
		My_addChar()
		nextToken = ADD_OP
	case '-':
		My_addChar()
		nextToken = SUB_OP
	case '*':
		My_addChar()
		nextToken = MULT_OP
	case '/':
		My_addChar()
		nextToken = DIV_OP
	}
}
