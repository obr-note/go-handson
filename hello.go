package main

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
)

type Mydata struct {
  Name string
  Mail string
  Tel string
}

func (m *Mydata) Str() string {
  return "<\"" + m.Name + "\" " + ", " + m.Tel + ">"
}

func main() {
  p := "https://tuyano-dummy-data.firebaseio.com/mydata.json"
  re, er := http.Get(p)
  if er != nil {
    panic(er)
  }
  defer re.Body.Close()

  s, er := ioutil.ReadAll(re.Body)
  if er != nil {
    panic(er)
  }

  var itms []Mydata
  er = json.Unmarshal(s, &itms)
  if er != nil {
    panic(er)
  }

  for i, im := range itms {
    println(i, im.Str())
  }
}
