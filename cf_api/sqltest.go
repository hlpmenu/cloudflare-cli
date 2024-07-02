package cfapi

import (
	"database/sql"
	"go-debug/output"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var testDB = &sqliteTestDB{}

type sqliteTestDB struct {
	DB     *sql.DB
	Active bool
}

const testDBPath = "test_db.db"

func IsValidSqlite(sql string) bool {
	var ok bool

	if testDB.DB == nil {
		testDB = &sqliteTestDB{}
		createDummyDB()
		testDB.Active = true
	}

	_, err := testDB.DB.Exec(sql)
	if err != nil {
		output.Infof("Error executing query: %s", sql)
		output.Errorf("Error executing query: %s", err)
		ok = false
	} else {
		ok = true
	}
	deleteDummyDB()
	return ok

}

// createDummyDB creates a new SQLite database for testing.
func createDummyDB() {
	file, err := os.Create(testDBPath)
	if err != nil {
		output.Errorf("Error creating test_db.db: %s", err)
		output.Exit("Exiting...")
	}
	file.Close()

	db, err := sql.Open("sqlite3", testDBPath)
	if err != nil {
		output.Errorf("Error opening test_db.db: %s", err)
		output.Exit("Exiting...")
	}

	testDB.DB = db
}

// deleteDummyDB closes the database connection and removes the test database file.
func deleteDummyDB() {
	err := testDB.DB.Close()
	if err != nil {
		output.Errorf("Error closing test_db.db: %s", err)
		output.Exit("Exiting...")
	}

	err = os.Remove("test_db.db")
	if err != nil {
		output.Errorf("Error removing test_db.db: %s", err)
		output.Exit("Exiting...")
	}

}

func parseSqlFile(path string) string {

	f, err := os.ReadFile(path)
	log.Printf("File: %s", f)
	if err != nil {
		if os.IsNotExist(err) {
			output.Errorf("Error: File not found")
		} else {
			output.Error("Error reading file")
		}
		output.Exit("Exiting...")
		return ""
	}
	query := strings.TrimSpace(string(f))
	log.Printf("Query: %s", query)

	if len(query) < 6 {
		output.Error("Invalid query")
		output.Exit("Please check your query and try again")
	}
	log.Printf("Query: %s", query)

	return query
}
