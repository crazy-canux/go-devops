package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlVersion(host string, port int, database, username, password string) (string, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
	conn, err := sql.Open("mysql", connString)
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
	err = conn.QueryRow(`select version()`).Scan(&version)
	if err == nil {
		log.Println("Get version failed.")
		return "", err
	}
	return version, nil
}
