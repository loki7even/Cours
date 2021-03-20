package sql

import (
	"database/sql"
	"strconv"
	"strings"

	"github.com/vosmith/pancake"
	_ "github.com/go-sql-driver/mysql"
)

type employees struct {
	id              int
	aicrew          int
	ground          int
	social_security int
	name            string
	first_name      string
	adress          string
}

func AddEmployees(aicrew int, ground int, social_security int, name string, first_name string, address string) {

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/aircraft")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `employees`(`aircrew`, `ground`, `social_security`, `name`, `first_name`, `address`) VALUES (?, ?, ?, ?, ?, ?)",
		aicrew, ground, social_security, name, first_name, address)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetEmployees(selector string, filter string) [][]string {

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/aircraft")
	if err != nil {
		panic(err.Error())
	}
	query := "SELECT "
	if selector != "" {
		query += selector
	} else {
		query += "* "
	}
	query += "FROM employees "
	if filter != "" {
		query += filter
	}

	query += ";"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var return_val [][]string
	var tag employees
	for selecte.Next() {
		selecte.Scan(&tag.id, &tag.adress, &tag.aicrew, &tag.ground, &tag.name, &tag.social_security, &tag.first_name)
		to_inject := []string{strconv.Itoa(tag.id), tag.adress, strconv.Itoa(tag.aicrew), strconv.Itoa(tag.ground),
			strconv.Itoa(tag.social_security), tag.name, tag.first_name}
		return_val = append(return_val, to_inject)
	}

	return return_val

}

func UpdateEmployees(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/aircraft")
	if err != nil {
		panic(err.Error())
	}

	query := "UPDATE `employees` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteEmployees(condition string) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/aircraft")
	if err != nil {
		panic(err.Error())
	}

	query := "DELETE FROM `employees` WHERE " + condition

	db.Query(query)

}

func GetPilotInfo() [][]string {
	
	pilote_id,_ := pancake.Strings(GetPilote("staff_id", ""))
	
	filter := "WHERE id IN ("+ strings.Join(pilote_id[:], ",") +")"
	return GetEmployees("", filter)
}