package main

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
)

var user string
var pass string
var host string
var port int

func Init(p map[string]interface{}) error {
  user = p["user"].(string)
  pass = p["pass"].(string)
  if p["host"] == nil {
    host = "localhost"
  } else {
    host = p["host"].(string)
  }
  if p["port"] == nil {
    port = 3306
  } else {
    port = int(p["port"].(float64))
  }
  return nil
}

func TablesNum(p map[string]interface{}) (string, error) {
  hostDSN := fmt.Sprintf("tcp(%s:%d)", host, port)
  dsn := user + ":" + pass + "@" + hostDSN + "/" + ""
  db, err := sql.Open("mysql", dsn)
  if err != nil {
    return "", err
  }
  defer db.Close()

  database := p["db"].(string)

  var tableNum string
  err = db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = ?", database).Scan(&tableNum)
  if err != nil {
    return "", err
  }
  return tableNum, nil
}

func ActiveConnections(p map[string]interface{}) (string, error) {
  hostDSN := fmt.Sprintf("tcp(%s:%d)", host, port)
  dsn := user + ":" + pass + "@" + hostDSN + "/" + ""
  db, err := sql.Open("mysql", dsn)
  if err != nil {
    return "", err
  }
  defer db.Close()

  var active string
  var val string
  err = db.QueryRow("show status where `variable_name` = 'Threads_connected'").Scan(&val,&active)
  if err != nil {
    return "", err
  }
  return active, nil
}
