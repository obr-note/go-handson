package main

import (
  "hello"
  "strconv"
  "fmt"
  "database/sql"
  _"github.com/mattn/go-sqlite3"
)

type Mydata struct {
  ID int
  Name string
  Mail string
  Age int
}

func (m *Mydata) Str() string {
  return "<\"" + strconv.Itoa(m.ID) + ":" + m.Name + "\" " + ", " + m.Mail + "," + strconv.Itoa(m.Age) + ">"
}

var qry string = "select * from mydata where name like ? or mail like ?"

func main() {
  con, er := sql.Open("sqlite3", "data.sqlite3")
  if er != nil {
    panic(er)
  }
  defer con.Close()

  for true {
    s := hello.Input("find")
    if s == "" {
      break
    }
    rs, er := con.Query(qry, "%"+s+"%", "%"+s+"%")
    if er != nil {
      panic(er)
    }
    for rs.Next() {
      var md Mydata
      er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
      if er != nil {
        panic(er)
      }
      fmt.Println(md.Str())
    }
  }
  fmt.Println("***end***")
}
