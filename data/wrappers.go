package data

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var GlobalConn SQLiteConn

type SQLiteConn struct {
	innerConn *gorm.DB
}

func (conn SQLiteConn) Query(sql string) *gorm.DB {
	log.Print(conn.innerConn)
	return conn.innerConn.Raw(sql)
}

func (conn SQLiteConn) Exec(sql string, args ...interface{}) *gorm.DB {
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
		log.Panic(err)
	}
	GlobalConn.innerConn = db
}

func init() {
	buildConn()
}
