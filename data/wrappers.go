package data

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var GlobalConn SQLiteConn

type SQLiteConn struct {
	innerConn *gorm.DB
}

func (conn SQLiteConn) initInstance(db *gorm.DB) {

	conn.innerConn = db
}

func (conn SQLiteConn) Query(sql string) *gorm.DB {
	return conn.innerConn.Raw(sql)
}

func (conn SQLiteConn) Exec(sql string, args ...interface{}) *gorm.DB {
	/*
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
	*/
	return conn.innerConn.Exec(sql, args)
}

func buildConn() {

	if GlobalConn.innerConn != nil {
		return
	}

	dbPath := os.Getenv("CACHE_DB")
	if dbPath == "" {
		dbPath = "./db.db"
	}

	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {

		panic(err)
	}
	GlobalConn.initInstance(db)
}

func init() {
	buildConn()
}
