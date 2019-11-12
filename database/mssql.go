package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func MssqlVersion(host string, port int, database, username, password string) (string, error) {
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;user id=%s;password=%s;encrypt=disable", host, port, database, username, password)
	conn, err := sql.Open("mssql", connString)
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
