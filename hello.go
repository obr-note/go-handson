package main

import (
  "fmt"
  "hello"
)

func main() {
  name := hello.Input("type your name")
  fmt.Println("Hello, " + name + "!!")
}
