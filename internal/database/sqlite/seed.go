package sqlite

import (
	"database/sql"
	"os"
)

func StartSchema(db *sql.DB, schemaFile string) error {
	schemaF, err := os.ReadFile(schemaFile)
	if err != nil {
		return err
	}
	if _, err = db.Exec(string(schemaF)); err != nil {
		return err
	}
	return nil
}
