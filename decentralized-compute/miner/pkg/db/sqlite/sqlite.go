package sqlite

import (
	"database/sql"
	"fmt"
	"os"
	"solo/pkg"

	_ "github.com/mattn/go-sqlite3"
)

const DRIVER = "sqlite3"

type SqliteDB struct {
	dbName string
	db     *sql.DB
}

func NewSqliteDB(dbName string) (*SqliteDB, error) {
	return &SqliteDB{dbName: dbName}, nil
}

func (s *SqliteDB) GetDB() *sql.DB {
	return s.db
}

func (s *SqliteDB) GetDBName() string {
	return s.dbName
}

func (s *SqliteDB) SetDBName(str string) {
	s.dbName = str
}

func (s *SqliteDB) Close() {
	s.db.Close()
}

func (s *SqliteDB) Connect() (*sql.DB, error) {
	dbDir := fmt.Sprintf("%s/db", pkg.CurrentDir())
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		if err = os.MkdirAll(dbDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create db directory: %v", err)
		}
	}

	db, err := sql.Open(DRIVER, fmt.Sprintf("%s/%s", dbDir, s.dbName))
	if err != nil {
		return nil, err
	}

	s.db = db
	return s.db, nil
}

/*
sqlStmt := `

	CREATE TABLE IF NOT EXISTS users (
	    id INTEGER PRIMARY KEY AUTOINCREMENT,
	    name TEXT,
	    age INTEGER
	);`
*/
func (s *SqliteDB) CreateTable(sqlStmt string) (interface{}, error) {
	return s.db.Exec(sqlStmt)
}

/*
*
"SELECT id, name, age FROM users"
*/
func (s *SqliteDB) Query(sqlStmt string) (*sql.Rows, error) {
	return s.db.Query(sqlStmt)
}

/*
*
"INSERT INTO users(name, age) VALUES(?, ?)"
*/
func (s *SqliteDB) Insert(sqlStmt string, args ...any) (sql.Result, error) {
	stmt, err := s.db.Prepare(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}
