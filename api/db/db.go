package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// sqlite database file
const dbFile string = "db.sqlite3"

// Directory with database migrations
const migrationsDir string = "api/migrations/"

// Open the SQL database and run migrations if necessary.
func Open() (*sql.DB, error) {
	// Open databse
	db, err := sql.Open("sqlite3", "file:"+dbFile+"?_foreign_keys=on")
	if err != nil {
		return nil, err
	}

	// Open migration files directory
	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		return nil, err
	}

	// Get all migration files by their IDs
	migrationIDs := make([]int, len(files))
	for i, file := range files {
		name := file.Name()
		nameNoExtension := name[:len(name)-4]
		id, err := strconv.Atoi(nameNoExtension)
		if err != nil {
			return nil, err
		}

		migrationIDs[i] = id
	}

	// Sort migrations by ID
	sort.Ints(migrationIDs)

	// Run initial migration tabel setup if not already setup
	_, migrationsTableExists := db.Query("SELECT * FROM migrations")
	if migrationsTableExists != nil {
		runMigration(db, 0)
	}

	// Get migrations that have already been run by querying the migrations table
	rows, err := db.Query("SELECT * FROM migrations")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	migrationsRun := make(map[int]bool)
	for rows.Next() {
		var migration int
		if err := rows.Scan(&migration); err != nil {
			return nil, err
		}
		migrationsRun[migration] = true
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// For every migration file, run it if its not in the migrations table (hasn't already been run)
	for _, id := range migrationIDs {
		if !migrationsRun[id] {
			if err := runMigration(db, id); err != nil {
				return nil, err
			}
		}
	}

	return db, nil
}

// Run a database migration.
func runMigration(db *sql.DB, id int) error {
	migrationFile := fmt.Sprintf("%v%d.sql", migrationsDir, id)

	migration, err := ioutil.ReadFile(migrationFile)
	if err != nil {
		fmt.Println("Unable to run migration", id)
		return err
	}

	fmt.Println("Running migration", id)
	if _, err := db.Exec(string(migration)); err != nil {
		return err
	}
	if _, err := db.Exec("INSERT INTO migrations VALUES (?)", id); err != nil {
		return err
	}
	return nil
}
