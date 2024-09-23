package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

var err error

func InitDB() {

	openDBConnection()

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()

	fmt.Println("DB OK!")
}

func openDBConnection() {
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		printErrorAndPanic(err, "Could not connet to database")
	}
}

func printErrorAndPanic(err error, message string) {
	fmt.Println(err)

	panic(message)
}

func createTables() {

	createTablesMap := getCreateTablesMap()

	for tableName, createTableSql := range createTablesMap {
		createTable(tableName, createTableSql)
	}
}

func createTable(tableName, createTableSql string) {
	_, err = DB.Exec(createTableSql)
	if err != nil {
		printErrorAndPanic(err, fmt.Sprintf("could not create %s table", tableName))
	}
}

func getCreateTablesMap() map[string]string {
	createTablesMap := map[string]string{
		"users": `
			CREATE TABLE  IF NOT EXISTS users(
				'id' INTEGER PRIMARY KEY AUTOINCREMENT,
				'name' TEXT NOT NULL UNIQUE,
				'email' TEXT NOT NULL UNIQUE,
				'height' DOUBLE(10,2) NOT NULL,
				'weight' DOUBLE(10,2) NOT NULL,
				'group' TINYINT 
			)
		`,
	}

	return createTablesMap
}
