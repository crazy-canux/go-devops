package main

import (
    "database/sql"
    "fmt"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    var server = "127.0.0.1"
    var port = 3306
    var database = "sandbox"
    var user = "username"
    var password = "password"
    connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, server, port, database)
    conn, err := sql.Open("mysql", connString)
    if err != nil {
        fmt.Println("connect error")
    }
    defer conn.Close()

    var version string
    err = conn.QueryRow(`select version()`).Scan(&version)
    if err == nil {
        fmt.Println(version)
    }
/*
    var rowsData []*Result
    for rows.Next() {
        var row = new(Result)
        rows.Scan(&row.Number, &row.AnalyzeType)
        rowsData = append(rowsData, row)
    }

    fmt.Println("Start to loop")
    for _, oneRow := range rowsData {
        fmt.Println(oneRow.Number, oneRow.AnalyzeType)
    }
    fmt.Println("loop finished.")
*/
}
