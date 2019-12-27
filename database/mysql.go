package database

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os/exec"

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

func PatchMysql(host string, port int, user, pw, db string, patch string) bool {
	var stdout, stderr bytes.Buffer
	h := fmt.Sprintf("-h%s", host)
	P := fmt.Sprintf("-P%d", port)
	u := fmt.Sprintf("-u%s", user)
	p := fmt.Sprintf("-p%s", pw)
	d := fmt.Sprintf("-D%s", db)
	query := fmt.Sprintf("source %s", patch)
	cmd := exec.Command("mysql", h, P, u, p, d, "-e", query)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	log.Printf("stdout: %s", stdout.String())
	log.Printf("stderr: %s", stderr.String())
	if err != nil {
		log.Printf("execute script error: %s", err.Error())
		return false
	}
	return true
}
