package db

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql" // ok
)

// DB db connection
type DB struct {
	db       *sql.DB
	host     string
	port     uint16
	username string
	password string
	dbname   string
}

// Open opens connection
// Now it's only suport mysql
func (db *DB) Open() error {
	var err error
	db.db, err = sql.Open("mysql", db.username+":"+db.password+"@tcp("+db.host+":"+strconv.Itoa(int(db.port))+")/"+db.dbname+"?charset=utf8mb4")
	err = db.db.Ping()
	return err
}

// Close closes connection
func (db *DB) Close() {
	db.db.Close()
}

// Query returns a map contains results
func (db *DB) Query(sql string) ([]map[string]string, error) {
	rows, err := db.db.Query(sql)
	if err != nil {
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	scanArgs := make([]interface{}, len(cols))
	values := make([]interface{}, len(cols))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	records := make([]map[string]string, 0)
	for rows.Next() {
		if err := rows.Scan(scanArgs...); err != nil {
			return nil, err
		}
		record := make(map[string]string)
		for i := range values {
			record[cols[i]] = string(values[i].([]byte))
		}
		records = append(records, record)
	}
	rows.Close()
	return records, nil

}

// Exec executes a sql and return the count of affected rows
func (db *DB) Exec(sql string) (int64, error) {
	tx, err := db.db.Begin()
	if err != nil {
		return 0, err
	}
	r, err := tx.Exec(sql)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	af, err := r.RowsAffected()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()
	return af, nil

}

// New returns a db connection
func New(host string, port uint16, username string, password string) (*DB, error) {
	db := new(DB)
	db.host = host
	db.port = port
	db.username = username
	db.password = password
	db.dbname = "gomeeting"
	return db, nil
}
