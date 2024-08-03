package database

import (
	"database/sql"
	"os"
	"reflect"

	"github.com/ddiogoo/broker/tree/master/key-manager-go/builder"
	_ "github.com/lib/pq"
)

type DatabaseManager struct {
	db        *sql.DB
	dbBuilder *builder.DatabaseQueryBuilder
}

// Close closes the connection with database.
func (k *DatabaseManager) Close() {
	err := k.db.Close()
	if err != nil {
		panic(err.Error())
	}
}

// Close closes the connection with database.
func (k *DatabaseManager) Ping() error {
	err := k.db.Ping()
	if err != nil {
		return err
	}
	return nil
}

// Table create all tables.
func (k *DatabaseManager) Table() error {
	k.dbBuilder.BuildCreateTable()
	for _, v := range k.dbBuilder.Tables {
		_, err := k.db.Exec(v)
		if err != nil {
			return err
		}
	}
	return nil
}

// NewDatabaseManager creates a new instance of DatabaseManager.
func NewDatabaseManager(types []reflect.Type) (*DatabaseManager, error) {
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

	return &DatabaseManager{
		db:        db,
		dbBuilder: builder.NewDatabaseQueryBuilder(types),
	}, nil
}
