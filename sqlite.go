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

  ids := hello.Input("update ID")
  id, _ := strconv.Atoi(ids)
  qry = "select * from mydata where id = ?"
  rw := con.QueryRow(qry, id)
  tgt := mydatafmRw(rw)
  fmt.Println(tgt.Str())
  f := hello.Input("delete it? (y/n)")
  if f == "y" {
    qry = "delete from mydata where id=?"
    con.Exec(qry, id)
  }
  showRecord(con)
}

func showRecord(con *sql.DB) {
  qry = "select * from mydata"
  rs, _ := con.Query(qry)
  for rs.Next() {
    fmt.Println(mydatafmRws(rs).Str())
  }
}

func mydatafmRws(rs *sql.Rows) *Mydata {
  var md Mydata
  er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
  if er != nil {
    panic(er)
  }
  return &md
}

func mydatafmRw(rs *sql.Row) *Mydata {
  var md Mydata
  er := rs.Scan(&md.ID, &md.Name, &md.Mail, &md.Age)
  if er != nil {
    panic(er)
  }
  return &md
}
