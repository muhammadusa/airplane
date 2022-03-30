package main

import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
host     = "localhost"
port     = 5432
user     = "muhammadsarvar"
password = "onam2575"
dbname   = "app"
)

var QUERY = `
	select 
		user_id,
		username
	from 
		users
` 

func main() {

connection := fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
	host, user, password, dbname, port,
)

db, err := sql.Open("postgres", connection)

defer db.Close()

if err != nil {
	panic(err)
}

var userId int
var username string

db.QueryRow(QUERY).Scan(&userId, &username)

fmt.Println(userId, username)

}