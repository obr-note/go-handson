package main

import (
  "io/ioutil"
  "os"
  "fmt"
)
  

func main() {
  rt := func(f *os.File) {
    s, er := ioutil.ReadAll(f)
    if er != nil {
      panic(er)
    }
    fmt.Println(string(s))
  }

  fn := "data.txt"

  f, er := os.OpenFile(fn, os.O_RDONLY, os.ModePerm)
  if er != nil {
    panic(er)
  }

  defer f.Close()

  fmt.Println("<< start >>")
  rt(f)
  fmt.Println("<< end >>")
}
