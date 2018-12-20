package main

import (
    "database/sql"
    "fmt"

    _ "github.com/zensqlmonitor/go-mssqldb"
)

type Result struct {
    Number uint32
    AnalyzeType string
}

func main() {
    var server = "127.0.0.1"
    var port = 1433
    var database = "sandbox"
    var user = "username"
    var password = "password"
    connString := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s;encrypt=disable", server, port, database, user, password)
    conn, err := sql.Open("mssql", connString)
    if err != nil {
        fmt.Println("connect error")
    }
    defer conn.Close()

    results, err := conn.Exec(`
select @@version
`)
    if err != nil {
        fmt.Println("query error")
    }
    // defer rows.Close()
    num, _ := results.RowsAffected()
    fmt.Println(num)
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
