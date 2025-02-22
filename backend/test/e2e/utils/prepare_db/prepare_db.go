package prepare_db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func PrepareDb(connectionString string) error {
	// Create mysql connection
	db, err := sql.Open(
		"mysql", connectionString,
	)

	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	err = db.Ping()
	if err != nil {
		return err
	}

	// Read all "up" migrations

	_, b, _, _ := runtime.Caller(0)
	currentDir := path.Join(path.Dir(b))

	migrationsDir := path.Join(currentDir, "../../../../config/db/migrations")

	err = filepath.WalkDir(migrationsDir, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if strings.Contains(path, ".up.sql") {
			// Execute
			println("Executing migration: " + path)

			fileContents, err := os.ReadFile(path)

			if err != nil {
				return err
			}

			_, err = db.Exec(string(fileContents))

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
