package data

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var GlobalConn SQLiteConn

type SQLiteConn struct {
	innerConn *sql.DB
}

func (conn SQLiteConn) initInstance(db *sql.DB) {

	conn.innerConn = db
}

func (conn SQLiteConn) Query(sql string) (*sql.Rows, error) {

	return conn.innerConn.Query(sql)
}

func (conn SQLiteConn) Exec(sql string, args ...interface{}) (int64, error) {
	stt, errStt := conn.innerConn.Prepare(sql)
	if errStt != nil {
		return 0, errStt
	}
	result, errExec := stt.Exec(args)
	if errExec != nil {
		return 0, errExec
	}
	rows, _ := result.RowsAffected()
	return rows, nil
}

func buildConn() {

	if GlobalConn.innerConn != nil {
		return
	}

	dbPath := os.Getenv("CACHE_DB")
	if dbPath == "" {
		dbPath = "./db.db"
	}

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {

		panic(err)
	}
	GlobalConn.initInstance(db)
}

func init() {
	buildConn()
}
