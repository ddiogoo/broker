package ctx

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

type KeyManagerDatabase struct {
	db *sql.DB
}

// Close closes the connection with database.
func (k *KeyManagerDatabase) Close() {
	err := k.db.Close()
	if err != nil {
		panic(err.Error())
	}
}

// Close closes the connection with database.
func (k *KeyManagerDatabase) Ping() error {
	err := k.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// NewKeyManagerDatabase creates a new instance of KeyManagerDatabase.
func NewKeyManagerDatabase() (*KeyManagerDatabase, error) {
	connStr := func() string {
		if os.Getenv("GIN_RUN_MODE") == "debug" {
			return os.Getenv("CONN_PGDB_DEBUG_DB")
		}
		return os.Getenv("CONN_PGDB_RELEASE_DB")
	}()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &KeyManagerDatabase{db: db}, nil
}
