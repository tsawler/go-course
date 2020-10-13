package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// DB holds a database connection
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

// ConnectMySQL creates database pool for mysql
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifetime)

	dbConn.SQL = d

	err = testDB(err, d)

	return dbConn, err
}

// testDB pings database
func testDB(err error, d *sql.DB) error {
	err = d.Ping()
	if err != nil {
		fmt.Println("Error!", err)
	} else {
		log.Println("*** Pinged database successfully! ***")
	}
	return err
}

// NewDatabase creates a new database for the app
func NewDatabase(dataSourceName string) (*sql.DB, error) {
	// create connection
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	// test ping
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
