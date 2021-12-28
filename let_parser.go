package main

import (
	"fmt"
	"os"
)

type astNode struct {
	parent   *astNode
	ttype    rune
	termsym  bool
	contents string
	children []*astNode
	env 	 []Binding
}

var parseTree astNode
var globalTree *astNode
var firstPass bool

func initTreeNd(nd *astNode, isterm bool) {
	nd.termsym = isterm
	nd.children = make([]*astNode, 0, 5)
	nd.ttype = tokenQueue[0].tokenType
	nd.contents = tokenQueue[0].tokenValue
	advanceToken()
}


func advanceToken() {
	//fmt.Println(tokenQueue[0].tokenValue)
	tokenQueue = tokenQueue[1:]
}


func checkToken(tokenQueue []Token, tokenVal string, errorMessage string) {
	
	if (tokenQueue[0].tokenValue == tokenVal) {
		advanceToken()
		
	} else {
		fmt.Println(errorMessage)
		os.Exit(1)
	}

}


func parseExp() astNode {
	root := astNode{}

	switch tokenQueue[0].tokenType {
	case 11: // LET
		initTreeNd(&root, false)
		child1 := parseExp()
		checkToken(tokenQueue, "=", "Error in parseExp: = sign not found in case let")
		child2 := parseExp()
		checkToken(tokenQueue, "in", "Error in parseExp: keyword in not found in case let")
		child3 := parseExp()
		root.children = append(root.children, &child1)
		root.children = append(root.children, &child2)
		root.children = append(root.children, &child3)

	case 22: // IF
		initTreeNd(&root, false)
		child1 := parseExp()
		checkToken(tokenQueue, "then", "Error in parseExp: then not found in case if")
		child2 := parseExp()
		checkToken(tokenQueue, "else", "Error in parseExp: else not found in case if")
		child3 := parseExp()
		root.children = append(root.children, &child1)
		root.children = append(root.children, &child2)
		root.children = append(root.children, &child3)

	case 23: // ISZERO
		initTreeNd(&root, false)
		checkToken(tokenQueue, "(", "Error in parseExp: lparen not found in case iszero")
		child1 := parseExp()
		checkToken(tokenQueue, ")", "Error in parseExp: rparen not found in case iszero")
		root.children = append(root.children, &child1)
		
	case 24: // MINUS
		initTreeNd(&root, false)
		checkToken(tokenQueue, "(", "Error in parseExp: lparen not found in case minus")
		child1 := parseExp()
		checkToken(tokenQueue, ",", "Error in parseExp: comma not found in case minus")
		child2 := parseExp()
		checkToken(tokenQueue, ")", "Error in parseExp: rparen not found in case minus")
		root.children = append(root.children, &child1)
		root.children = append(root.children, &child2)

	default: // LEAVES
		initTreeNd(&root, true)
	}

	return root
}


func printTree(nd astNode) {
	nd.printTreeWork(0)
}


func (nd astNode) printTreeWork(indentLevel int)   {
	outString := ""
	for i := 0; i < indentLevel; i++ {
		outString += "    "
	}

	switch nd.ttype {
	case 11: // LET
		fmt.Printf("%s", outString)
		if nd.env == nil && firstPass == false {
			fmt.Println("LetExp(")
		} else {
			fmt.Println("LetExp(", "e = ", nd.env)
		}
		outString += "    "
		fmt.Printf("%s", outString)
		fmt.Printf("%q", nd.children[0].contents)
		fmt.Print(",")
		fmt.Println()
		
	case 24: // MINUS
		fmt.Printf("%s", outString)
		if nd.env == nil {
			fmt.Println("DiffExp(")
		} else {
			fmt.Println("DiffExp(", "e = ", nd.env)
		}
	case 22: // IF STMT
		fmt.Printf("%s", outString)
		if nd.env == nil {
			fmt.Println("IfExp(")
		} else {
			fmt.Println("IfExp(", "e = ", nd.env)
		}
		
	case 23: // ISZERO
		fmt.Printf("%s", outString)
		if nd.env == nil {
			fmt.Println("IsZeroExp(")
		} else {
			fmt.Println("IsZeroExp(", "e = ", nd.env)
		}
		
	case 13: // CONST
		fmt.Printf("%s", outString)
		fmt.Println("ConstExp(")
		fmt.Printf("%s", outString)
		fmt.Print("    ")
		fmt.Println(nd.contents)
		fmt.Printf("%s", outString)
		fmt.Println(")")
		
	case 12: // VAR
		fmt.Printf("%s", outString)
		fmt.Println("VarExp(")
		fmt.Printf("%s", outString)
		fmt.Print("    ")
		fmt.Printf("%q", nd.contents)
		fmt.Println()
		fmt.Printf("%s", outString)
		fmt.Println(")")
	}
	
	if nd.termsym == false && len(nd.children) > 0 {

		theLoop := len(nd.children)
		for i := 0; i < theLoop; i++ {
			if nd.ttype == 11 {
				nd.children[i+1].printTreeWork(indentLevel+1)
				theLoop--
			} else {
				nd.children[i].printTreeWork(indentLevel+1)
			}
		}
		if nd.ttype != 11 {
			fmt.Printf("%s", outString)
			fmt.Println(")")
		} else {
			outString = outString[3:]
			fmt.Printf("%s", outString)
			fmt.Println(")")
		}	
	}
}
	

func Parser() {

	parseTree := parseExp()
	printTree(parseTree)
	Evaluator(parseTree)
	printTree(parseTree)
	Evaluator(parseTree)
}