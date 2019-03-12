/*
 * File: database.go
 * Project: Moneway Intern Assesment
 * File Created: Monday, 11th March 2019 9:17:12 AM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.0
 * Brief: Creating a MySQL Database Package with PDO.
 * -----
 * Last Modified: Monday, 11th March 2019 9:24:05 AM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/net/context"
)

type Config struct {
	Database database
}

type database struct {
	Server   string
	Port     string
	Database string
	User     string
	Password string
}

var db *sql.DB

/**
 * @brief This Initialize's the Database Instance.
 *
 * @param
 *
 * @return void
 */
func Init() {
	var config Config
	if _, err := toml.DecodeFile("../config/config.toml", &config); err != nil {
		fmt.Println(err)
	}
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Database.User, config.Database.Password, config.Database.Server, config.Database.Port, config.Database.Database)

	dbConn, err := sql.Open("mysql", connString)
	checkError(err)
	db = dbConn

	fmt.Println("Database Is Connected")
}

/**
 * @brief This Inserts Into The Database
 *
 * @param ctx context.Context //Context
 * @param table string //The table you want to access.
 * @param columns []string //The Columns you want to insert the data.
 * @param values []string //The data to inserted.
 *
 * @return bool
 */
func Insert(ctx context.Context, table string, columns []string, values []string) bool {
	success := true
	sqlStmt := "INSERT INTO " + table + "("
	questionMarks := "("
	count := 0
	for _, column := range columns {
		if count != len(columns)-1 {
			sqlStmt += "" + column + ", "
			questionMarks += "?, "
		} else {
			sqlStmt += "" + column + ") VALUES"
			questionMarks += "?)"
		}
		count++
	}
	sqlStmt += questionMarks

	args := make([]interface{}, len(values))
	for i := range values {
		args[i] = values[i]
	}

	_, err := db.ExecContext(ctx, sqlStmt, args...)
	success = checkError(err)

	return success
}

/**
 * @brief This Gets Data From The Database
 *
 * @param ctx context.Context //Context
 * @param table string //The table you want to access
 * @param condition []string //The Condition that must hold.
 *
 * @return string
 */
func Select(ctx context.Context, table string, condition []string) string {

	var id = condition[1]
	sqlStmt := "select " + condition[2] + " from " + table + " where " + condition[0] + " = ?"
	fmt.Println(sqlStmt)

	var column string
	stmt, err := db.PrepareContext(ctx, sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	err = stmt.QueryRow(id).Scan(&column)
	checkError(err)
	fmt.Println("DB: ", column)

	return column
}

/**
 * @brief This Update A Column in A Table Data From The Database
 *
 * @param ctx context.Context //Context
 * @param table string //The table you want to access
 * @param params []string //The Columns you want to insert the data
 *
 * @return bool
 */
func Update(ctx context.Context, table string, params []string) bool {
	success := true
	sqlStmt := "UPDATE " + table + " SET " + params[2] + " = ? WHERE " + params[0] + " = ?"

	_, err := db.ExecContext(ctx, sqlStmt, params[3], params[1])
	success = checkError(err)

	return success
}

/**
 * @brief This Checks if there is an error
 *
 * @param error err //The error
 *
 * @return bool
 */
func checkError(err error) bool {
	success := true
	if err != nil {
		success = false
		log.Fatal(err)
	}
	return success
}

// This is for testing out the various functions
// func main() {
// 	ctx, stop := context.WithCancel(context.Background())
// 	defer stop()
// 	Init()
// 	table := "account"

// 	columns := []string{"firstname", "lastname", "balance", "currency"}
// 	values := []string{"Tassie", "Antwi-Donkor", "5000", "USD"}
// 	Insert(ctx, table, columns, values)

// 	conditions := []string{"account_id", "1"}
// 	Select(ctx, table, conditions)

// 	params := []string{"account_id", "1", "balance", "2690.90"}
// 	Update(ctx, table, params)
// }
