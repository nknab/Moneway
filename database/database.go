/*
 * File: database.go
 * Project: Moneway Go Developper Intern Challenge
 * File Created: Friday, 15th March 2019 7:48:11 PM
 * Author: nknab
 * Email: kojo.anyinam-boateng@outlook.com
 * Version: 1.1
 * Brief: This is the database class to handle all database queries.
 * -----
 * Last Modified: Sunday, 17th March 2019 8:10:57 PM
 * Modified By: nknab
 * -----
 * Copyright ©2019 nknab
 */

package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	ut "github.com/nknab/Moneway/util"
	"golang.org/x/net/context"
)

/**
 * @brief This Initialize's the Database Instance.
 *
 * @param filePath string //File path to the env file.
 *
 * @return void
 */
func Init(filePath string) *sql.DB {
	err := godotenv.Load(filePath)
	ut.CheckError(err, "Error loading config.env file")
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_SERVER"), os.Getenv("DB_PORT"), os.Getenv("DB"))

	dbConn, err := sql.Open("mysql", connString)
	ut.CheckError(err, "Can not connect to Database")

	return dbConn
}

/**
 * @brief This Inserts Into The Database
 *
 * @param ctx context.Context //Context
 * @param table string //The table you want to access.
 * @param columns []string //The Columns you want to insert the data.
 * @param values []string //The data to inserted.
 * @param db *sql.DB //A reference to a database instance.
 *
 * @return int32
 * @return bool
 */
func Insert(ctx context.Context, table string, columns []string, values []string, db *sql.DB) (int32, bool) {
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
	success = ut.CheckError(err, "Insert Query Could Not be Executed")

	var id int32
	//Getting the ID of the last insertion
	if success {
		sqlStmt = "select max(transaction_id) from " + table
		stmt, err := db.PrepareContext(ctx, sqlStmt)
		ut.CheckError(err, "Binding failed")
		err = stmt.QueryRow().Scan(&id)
		ut.CheckError(err, "Select Query Could Not be Executed")
	} else {
		id = -1
	}

	return id, success
}

/**
 * @brief This Gets Data From The Database
 *
 * @param ctx context.Context //Context
 * @param table string //The table you want to access
 * @param condition []string //The Condition that must hold.
 * @param db *sql.DB //A reference to a database instance.
 *
 * @return string
 */
func Select(ctx context.Context, table string, condition []string, db *sql.DB) string {

	var id = condition[1]
	sqlStmt := "select " + condition[2] + " from " + table + " where " + condition[0] + " = ?"

	var column string
	stmt, err := db.PrepareContext(ctx, sqlStmt)
	ut.CheckError(err, "Binding failed")
	err = stmt.QueryRow(id).Scan(&column)
	ut.CheckError(err, "Select Query Could Not be Executed")

	return column
}

/**
 * @brief This Update A Column in A Table Data From The Database
 *
 * @param ctx context.Context //Context
 * @param table string //The table you want to access
 * @param params []string //The Columns you want to insert the data
 * @param db *sql.DB //A reference to a database instance.
 *
 * @return bool
 */
func Update(ctx context.Context, table string, params []string, db *sql.DB) bool {
	success := true
	sqlStmt := "UPDATE " + table + " SET " + params[2] + " = ? WHERE " + params[0] + " = ?"

	_, err := db.ExecContext(ctx, sqlStmt, params[3], params[1])
	success = ut.CheckError(err, "Update Query Could Not be Executed")

	return success
}
