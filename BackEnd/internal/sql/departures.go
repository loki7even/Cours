package sql

import (
	"database/sql"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type departures struct {
	id          int
	id_flight   int
	date        time.Time
	pilote      int
	copilote    int
	aircrew     string
	free_places int
	occupied    int
}

func AddDepartures(id_flight int, date time.Time, pilote int, copilote int, aircrew string, free_places int, occupied int) {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `departures`(`id_fligth`, `date`, `pilote`, `copilote`, `aircrew`, `free_places`, `occupied`, `ticket_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		id_flight, date, pilote, copilote, aircrew, free_places, occupied)

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}

	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func GetDepartures(selector string, filter string) [][]string {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")
	if err != nil {
		panic(err.Error())
	}
	query := "SELECT "
	if selector != "" {
		query += selector
	} else {
		query += "* "
	}
	query += "FROM departures "
	if filter != "" {
		query += " WHERE `id` IN (" + filter + ")"
	}

	query += ";"

	selecte, err := db.Query(query)

	if err != nil {
		panic(err.Error())
	}

	var return_val [][]string
	var tag departures
	for selecte.Next() {
		selecte.Scan(&tag.id, &tag.id_flight, &tag.date, &tag.aircrew, &tag.copilote, &tag.pilote, &tag.free_places,
			&tag.occupied)
		to_inject := []string{strconv.Itoa(tag.id), strconv.Itoa(tag.id_flight), tag.date.Format(time.UnixDate),
			tag.aircrew, strconv.Itoa(tag.copilote), strconv.Itoa(tag.pilote), strconv.Itoa(tag.free_places),
			strconv.Itoa(tag.occupied)}
		return_val = append(return_val, to_inject)
	}

	return return_val

}

func UpdateDepartus(column string, new_value string, condition string) {

	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")
	if err != nil {
		panic(err.Error())
	}

	query := "UPDATE `departures` SET " + column + " " + new_value + " WHERE " + condition
	db.Query(query)

}

func DeleteDepartus(condition string) {
	db, err := sql.Open("mysql", "root:passwd@tcp(172.21.0.2:3306)/aircraft")
	if err != nil {
		panic(err.Error())
	}

	query := "DELETE FROM `departures` WHERE " + condition

	db.Query(query)

}
