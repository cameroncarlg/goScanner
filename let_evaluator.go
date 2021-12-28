package main

import (
	"fmt"
	"strconv"
)

var isZero bool

// Binding represents a var/value pair
type Binding struct {
	varname string
	value rune
}


func lookup(varname string, e []Binding) rune {
	for i := range e {
		if e[i].varname == varname {
			return e[i].value
		}
	}
	return 0
}

func evaluate(localRoot *astNode, e []Binding) rune {
	localRoot.env = e
	switch localRoot.ttype {
	case 11: // LET
		 varname := localRoot.children[0].contents
		 exp1Val := evaluate(localRoot.children[1], e)
		 var newBindingList []Binding = []Binding{Binding{varname, exp1Val}}
		 e = append(newBindingList, e...)
		 return evaluate(localRoot.children[2], e)
	case 24: // MINUS
		exp1 := evaluate(localRoot.children[0], e)
		exp2 := evaluate(localRoot.children[1], e)
		return (exp1 - exp2)
	case 22: // IF STMT
		exp1 := evaluate(localRoot.children[0], e)
		if isZero || exp1 == 0 {
			return evaluate(localRoot.children[1], e)
		} else {
			return evaluate(localRoot.children[2], e)
		}
	case 23: // IS ZERO
		exp1 := evaluate(localRoot.children[0], e)
		if exp1 == 0 {
			isZero = true
		}
		return exp1
	case 12: // IDENTIFIER
		return lookup(localRoot.contents, e)
	case 13: // INTEGER
		temp, err := strconv.Atoi(localRoot.contents)
		_ = err
		return int32(temp)
	default:
		return 0
	}
}


func Evaluator(nd astNode) {
	e := []Binding{}
	fmt.Println()
	result := (evaluate(&nd, e))
	fmt.Println("Result: ", result)
	fmt.Println()
	firstPass = true
}
