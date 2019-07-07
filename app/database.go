package main

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Number struct {
	id       string
	unixtime int
	val      string
}

func PutRand(value string, db *sql.DB) (string, int) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", 1
	}
	_, err = db.Exec("insert into table_numbers (id, val, unixtime) values ($1, $2, $3)",
		id.String(), value, int32(time.Now().Unix()))
	if err != nil {
		return "", 1
	}
	return id.String(), 0
}

func GetRand(id string, db *sql.DB) (Number, int) {
	dt := Number{}
	row := db.QueryRow("select * from table_numbers where id = $1", id)

	err := row.Scan(&dt.id, &dt.val, &dt.unixtime)

	if err != nil {
		return Number{}, 1
	}
	return dt, 0
}
