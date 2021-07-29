package main

import (
  "io/ioutil"
  "fmt"
)
  

func main() {
  fs, er := ioutil.ReadDir(".")
  if er != nil {
    panic(er)
  }

  for _, f := range fs {
    fmt.Println(f.Name(), "(", f.Size(), ")")
  }
}
