package main

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
)

var user string
var pass string
var host string
var port string

func Init(p map[string]interface{}) {
  user = p["user"].(string)
  pass = p["pass"].(string)
  if p["host"] == nil {
    host = "localhost"
  } else {
    host = p["host"].(string)
  }
  if p["port"] == nil {
    port = "3306"
  } else {
    port = p["port"].(string)
  }
}

func TablesNum(p map[string]interface{}) string {
  hostDSN := fmt.Sprintf("tcp(%s:%s)", host, port)
  dsn := user + ":" + pass + "@" + hostDSN + "/" + ""
  db, err := sql.Open("mysql", dsn)
  database := p["db"].(string)
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()

  var tableNum string
  db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = ?", database).Scan(&tableNum)
  return tableNum
}

func ActiveConnection(p map[string]interface{}) string {
  hostDSN := fmt.Sprintf("tcp(%s:%s)", host, port)
  dsn := user + ":" + pass + "@" + hostDSN + "/" + ""
  db, err := sql.Open("mysql", dsn)
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()

  var active string
  var val string
  db.QueryRow("show status where `variable_name` = 'Threads_connected'").Scan(&val,&active)
  return active
}
