package main

import (
  _ "fmt"
  _ "log"
  _ "strings"
)

type token struct {
  kind  string 
  value string 
}

func tokenizer(input string) []token {
  input += "\n"

  current := 0 

  tokens := []token{}

  for current < len([]rune(input)) {
    char := string([]rune(input)[current])

    if char == "(" {
      tokens = append(tokens, token{
        kind: "paren",
        value:  "(",
      })

      current++

      continue
    }

    if char == ")" {
      tokens = append(tokens, token{
        kind: "paren",
        value:  ")",
      })
      current ++
      continue
    }


    if char == " " {
      current++
      continue
    }


    if isNumber(char) {
      value := ""

      
      for isNumber(char) {
        value += char 
        current++ 
        char = string([]rune(input)[current])
      }

      tokens = append(tokens, token{
        kind: "number",
        value: value,
      })

      continue

    }


    if isLetter(char) {
      value := ""

      for isLetter(char) {
        value += char 
        current++ 
        char = string([]rune(input)[current])
      }

      tokens = append(tokens, token{
        kind: "name",
        value: value,
      })
      continue
    }
    break
  }
  return tokens
}


// isNumber accepts a string and will check to see whether or not what has been
// passed through is between the range of 0 - 9.
func isNumber(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	if n >= '0' && n <= '9' {
		return true
	}
	return false
}

// isLetter works in a similar way to isNumber, but checks the range for a
// letter in the range of a - z.
func isLetter(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	if n >= 'a' && n <= 'z' {
		return true
	}
	return false
}


 type node struct {
	kind       string
	value      string
	name       string
	callee     *node
	expression *node
	body       []node
	params     []node
	arguments  *[]node
	context    *[]node
}

// Type `ast` is just another alias type. I find this makes part of the code
// more readable, as you'll come to see that there are a ton of references to
// `node`.
type ast node

// This is the counter variable that we'll use for parsing.
var pc int

// This variable will store our slice of `token`s inside of it.
var pt []token
