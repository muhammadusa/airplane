package main

import(
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "muhammadsarvar"
	password = "onam2575"
	dbname = "air"
)

var QUERY = `
select
	airplane_id,
	terminal ,
	flight,
	to_char(flytime, 'HH24:MI'),	
	togo,
	gate,
	to_char(landtime, 'HH24:MI'),
	remark
from 
	air`

// func main() {
// 	connection := fmt.Sprintf(
// 		"host=%s user=%s password=%s dbname=%s port=%d",
// 		host, user, password, dbname, port,
// 	)

// 	db, err := sql.Open("postgres", connection)

// 	defer db.Close()

// 	if err != nil {
// 		panic(err)
// 	}

// 	var terminal int 
// 	var airplane_id int 
// 	var flight string
// 	var flytime string
// 	var togo string
// 	var gate int 
// 	var landtime string
// 	var remark string

// 	db.QueryRow(QUERY).Scan( &terminal, &airplane_id, &flight, &flytime, &togo, &gate, &landtime, &remark)

// 	fmt.Println(airplane_id, terminal, flight, flytime, togo, gate, landtime, remark)

// }

type Air struct {
	Terminal int 
	Airplane_id int 
	Flight string
	Flytime string
	Togo string
	Gate int 
	Landtime string
	Remark string
}

func GetAir() []Air {
	connection := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		host, user, password, dbname, port,
	)

	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, _ := db.Query(QUERY)

	defer rows.Close()

	var airs []Air 

	for rows.Next(){
		var air Air 

		rows.Scan( &air.Airplane_id, &air.Terminal, &air.Flight, &air.Flytime, &air.Togo, &air.Gate, &air.Landtime, &air.Remark)

		airs = append(airs, air)
	}

	// db.QueryRow(QUERY).Scan( &terminal, &airplane_id, &flight, &flytime, &togo, &gate, &landtime, &remark)

	// fmt.Println(airplane_id, terminal, flight, flytime, togo, gate, landtime, remark)
	return airs
}