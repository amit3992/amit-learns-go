package main

import "fmt"

func isOpenChar(c rune) bool {
  if c == '{' || c == '(' || c == '[' {
    return true;
  } else {
    return false;
  }
}

func isCloseChar(c rune) bool {
  if c == '}' || c == ')' || c == ']' {
    return true;
  } else {
    return false;
  }
}

func isBalanced(input string) bool {

  count := 0
  for _,c := range input {

    if isOpenChar(c) {
      count += 1
    }

    if isCloseChar(c) {
      count -= 1
    }

  }

  if count != 0 {
    return false;
  } else {
    return true;
  }

}


func main() {

  if isBalanced("{}(") {
    fmt.Println("Is Balanced")
  } else {
    fmt.Println("Not balanced")
  }
}
