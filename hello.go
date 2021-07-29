package main

import (
  "net/http"
  "io/ioutil"
  "fmt"
  "encoding/json"
)
  

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

  var data []interface{}
  er = json.Unmarshal(s, &data)
  if er != nil {
    panic(er)
  }

  for i, im := range data {
    m := im.(map[string]interface{})
    fmt.Println(i, m["name"].(string), m["mail"].(string), m["tel"].(string))
  }
}
