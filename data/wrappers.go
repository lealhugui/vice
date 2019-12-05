package data

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var GlobalConn SQLiteConn

type SQLiteConn struct {
	DB *gorm.DB
}

func (conn SQLiteConn) Query(sql string) *gorm.DB {
	return conn.DB.Raw(sql)
}

func (conn SQLiteConn) Exec(sql string, args ...interface{}) *gorm.DB {
	return conn.DB.Exec(sql, args...)
}

func buildConn() {

	if GlobalConn.DB != nil {
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
	GlobalConn.DB = db
}

func init() {
	buildConn()
}
