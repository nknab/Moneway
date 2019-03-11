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

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
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

	db, _ = sql.Open("mysql", connString)

	defer db.Close()
	fmt.Println("Go MySQL Tutorial")
}

/**
 * @brief This Inserts Into The Database
 *
 * @param table string //The table you want to access.
 * @param columns []string //The Columns you want to insert the data.
 * @param values []string //The data to inserted.
 *
 * @return bool
 */
func Insert(table string, columns []string, values []string) bool {
	success := true
	sql := "INSERT INTO " + table + "("
	questionMarks := "("
	count := 0
	for _, column := range columns {
		if count != len(columns)-1 {
			sql += "" + column + ", "
			questionMarks += "?, "
		} else {
			sql += "" + column + ") VALUES"
			questionMarks += "?)"
		}
		count++
	}

	sql += questionMarks

	tx, err := db.Begin()
	if err != nil {
		success = false
		fmt.Println(err)
	}
	defer tx.Rollback()
	fmt.Println(sql)
	stmt, err := tx.Prepare(sql)
	if err != nil {
		success = false
		fmt.Println(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(values[:])
	if err != nil {
		success = false
		fmt.Println(err)
	}

	return success
}

/**
 * @brief This Gets Data From The Database
 *
 * @param table string //The table you want to access
 * @param condition []string //The Condition that must hold.
 *
 * @return string
 */
func Select(table string, condition []string) string {
	sql := "SELECT * FROM " + table + " WHERE `" + condition[0] + "` = ?"

	fmt.Println(sql)

	data, err := db.Query(sql, 1)
	if err != nil {
		success = false
		fmt.Println(err)
	}
	defer data.Close()

	return data
}

/**
 * @brief This Update A Column in A Table Data From The Database
 *
 * @param table string //The table you want to access
 * @param columns []string //The Columns you want to insert the data
 *
 * @return bool
 */
func Update(table string, condition []string) bool {
	success := true
	// sql := "UPDATE " + table + " SET `" + condition[0] + "` = ? WHERE " + condition[2] = " ?"

	// fmt.Println(sql)

	// data, err := db.Query(sql, 1)
	// if err != nil {
	// 	success = false
	// 	fmt.Println(err)
	// }
	// defer data.Close()
	// fmt.Println(data)

	return success
}

//func main(){
//	Init()
//	table := "account"
//	columns := []string{"firstname", "lastname", "balance", "currency"}
//	values := []string{"Tassie", "Antwi-Donkor", "5000", "USD"}
//
//	Insert(table, columns, values)
//
//	conditions := []string{"account_id", "1"}
//	Select(table, conditions)
//}
