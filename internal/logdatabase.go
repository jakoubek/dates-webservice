package internal

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

type Logdatabase struct {
	databasePath string
	db           *sql.DB
}

type LogdatabaseConfig func(ldb *Logdatabase)

func NewLogdatabase(opts ...LogdatabaseConfig) *Logdatabase {

	ldb := Logdatabase{}

	for _, opt := range opts {
		opt(&ldb)
	}

	err := ldb.OpenDatabase()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database connection ok")
	}
	ldb.createSchema()

	return &ldb

}

func WithDatabasePath(path string) LogdatabaseConfig {
	return func(ldb *Logdatabase) {
		ldb.databasePath = path
	}
}

func (ldb *Logdatabase) CreateRequest(request string, lang string, format string) {

	stmt, err := ldb.db.Prepare("INSERT INTO request(request, lang, format) VALUES(?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(request, lang, format)
	if err != nil {
		log.Println(err)
	}

}

func (ldb *Logdatabase) createSchema() error {

	ddl := `CREATE TABLE IF NOT EXISTS request (
    id INTEGER PRIMARY KEY,
    request TEXT,
    lang TEXT,
    format TEXT,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP
);`
	_, err := ldb.db.Exec(ddl)
	if err != nil {
		return fmt.Errorf("error creating database schema: %s", err)
	}
	return nil
}

func (ldb *Logdatabase) OpenDatabase() error {
	dataSourceName := fmt.Sprintf("file:%s?journal_mode=wal", ldb.databasePath)
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return fmt.Errorf("error opening database: %s", err)
	}
	ldb.db = db
	return nil
}

func (ldb *Logdatabase) CloseDatabase() error {
	err := ldb.db.Close()
	if err != nil {
		return fmt.Errorf("error closing database: %s", err)
	}
	return nil
}
