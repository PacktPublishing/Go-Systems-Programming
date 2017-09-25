package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"text/template"
)

func main() {
	var username string
	var password string

	arguments := os.Args
	if len(arguments) == 3 {
		username = arguments[1]
		password = arguments[2]
	} else {
		fmt.Println("programName Username Password!")
		os.Exit(100)
	}

	connectString := username + ":" + password + "@unix(/tmp/mysql.sock)/information_schema"
	db, err := sql.Open("mysql", connectString)

	rows, err := db.Query("SELECT DISTINCT(TABLE_SCHEMA) FROM TABLES;")
	if err != nil {
		fmt.Println(err)
		os.Exit(100)
	}

	var DATABASES []string
	for rows.Next() {
		var databaseName string
		err := rows.Scan(&databaseName)
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}
		DATABASES = append(DATABASES, databaseName)
	}
	db.Close()

	t := template.Must(template.New("t1").Parse(`
	{{range $k := .}} {{ printf "\tDatabase Name: %s" $k}}
	{{end}}
	`))
	t.Execute(os.Stdout, DATABASES)
	fmt.Println()
}
