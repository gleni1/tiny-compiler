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
