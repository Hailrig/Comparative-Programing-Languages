package main //By Lucas Harvey

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Lucas' Recursive parser")
	for {
		fmt.Println("Enter the next string to be parsed, or 'end' to end the program: ") //aquire input string

		stringReader := bufio.NewReader(os.Stdin)
		input, _ := stringReader.ReadString('\n')
		input = strings.TrimRight(input, "\r\n")
		if input == "end" { //check for ending the program
			break
		}

		inputArray, valid := scanner(input) //scan for lexemes, return lexeme array and error status
		if valid == false {                 //if not lexically valid, end
			fmt.Println(input, "contains lexical units which are not lexemes (valid)")
			continue
		}

		inputArray = append(inputArray, "Sentinal")
		value, pointer := E(inputArray, 0)                 //Check for Syntactic validity
		if value == true && pointer == len(inputArray)-1 { // if valid, report as such
			fmt.Println(input, "is an expression")
		} else { //if invalid, report as such
			fmt.Println(input, "is not an expression")
		}
	}

}

func scanner(input string) ([]string, bool) { //scan for lexemes, return lexeme array and error status
	inputArray := strings.Fields(input) //purge whitespaces, store lexemes in array

	for i := 0; i < len(inputArray); i++ { //check every element of the array to verify that they are valid lexemes
		switch inputArray[i] {
		case
			"1",
			"2",
			"3",
			"4",
			"5",
			"6",
			"7",
			"8",
			"9",
			"-",
			"*",
			"+":
			continue
		}
		return inputArray, false //report that they are not
	}
	return inputArray, true //report that they are
}

func E(inputArray []string, pointer int) (bool, int) { //Check to see if E has an appropriate next move, and perform it if so
	fmt.Println("E")
	storedPointer := pointer
	var value bool
	value, pointer = T(inputArray, pointer) //try if T is an appropriate next move
	if value == true {
		if inputArray[pointer] == "+" { //if T worked, check if the next value is a +
			pointer++
			value, pointer = E(inputArray, pointer) //if the next value was a + and T worked, try E
			if value == true {
				return true, pointer //return affimitively with new pointer
			} else {
				pointer--
				return true, pointer //return affimitively with new pointer
			}
		} else {
			return true, pointer //return affimitively with new pointer
		}
	} else {
		return false, storedPointer //return negatively with old pointer
	}
}

func T(inputArray []string, pointer int) (bool, int) { //Check to see if T has an appropriate next move, and perform it if so
	fmt.Println("T")
	storedPointer := pointer
	var value bool
	value, pointer = N(inputArray, pointer) //try if N is an appropriate next move
	if value == true {
		if inputArray[pointer] == "*" { //if N worked, check if the next value is a *
			pointer++
			value, pointer = T(inputArray, pointer) //if the next value was a * and N worked, try T
			if value == true {
				return true, pointer //return affimitively with new pointer
			} else {
				pointer--
				return true, pointer //return affimitively with new pointer
			}
		} else {
			return true, pointer //return affimitively with new pointer
		}
	} else {
		return false, storedPointer //return negatively with old pointer
	}
}

func N(inputArray []string, pointer int) (bool, int) { //Check to see if N has an appropriate next move, and perform it if so
	fmt.Println("N")
	storedPointer := pointer
	var value bool
	if inputArray[pointer] == "-" { //check to see if the second grammer rule of N is appropriate, recursively call D if so
		pointer++
		value, pointer = D(inputArray, pointer)
		if value {
			return true, pointer //return affimitively with new pointer
		} else {
			return false, storedPointer //return negatively with old pointer
		}
	} else { //try the first grammer rule of N, calling D in the process
		value, pointer = D(inputArray, pointer)
		if value {
			return true, pointer //return affimitively with new pointer
		} else {
			return false, storedPointer //return negatively with old pointer
		}
	}
}

func D(inputArray []string, pointer int) (bool, int) { //check to see if D is an appropriate next move, and perform it if it is
	fmt.Println("D")
	storedPointer := pointer
	switch inputArray[pointer] { //check if we can apply the first rule of the nonterminal
	case
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9":
		pointer++
		fmt.Println("true")
		return true, pointer //return affirmitively with new pointer
	}
	return false, storedPointer //return negatively with old pointer
}
