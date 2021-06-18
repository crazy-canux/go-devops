package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
)

func MssqlVersion(host string, port int, database, username, password string) (string, error) {
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s;encrypt=disable", host, port, database, username, password)
	//conn, err := sql.Open("mssql", connString)
	conn, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Println("Parameter error.")
		return "", err
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		log.Println("Connection failed.")
		return "", err
	}

	var version string
	err = conn.QueryRow(`select @@version`).Scan(&version)
	if err != nil {
		log.Println("Get version failed.")
		return "", err
	}

	return version, nil
}

func PatchMssql(host string, port int, user, pw, db string, patch string) bool {
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s;encrypt=disable", host, port, db, user, pw)
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Printf("database %s parameter error: %s", host, err.Error())
		return false
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		log.Printf("Connection error: %s", err.Error())
		return false
	}

	fBytes, err := ioutil.ReadFile(patch)
	if err != nil {
		log.Printf("open db script error: %s", err.Error())
		return false
	}

	sqlSlice1 := strings.Split(string(fBytes), "GO")
	sqlSlice2 := strings.Split(strings.Join(sqlSlice1, ""), "go")
	query := strings.Join(sqlSlice2, "")

	stmt, err := conn.Prepare(query)
	if err != nil {
		log.Printf("perpare script error: %s", err.Error())
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		log.Printf("execute script error: %s", err.Error())
		return false
	}
	return true
}
