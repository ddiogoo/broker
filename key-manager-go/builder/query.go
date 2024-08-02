package builder

import (
	"fmt"
	"reflect"
)

type DatabaseQueryBuilder struct {
	types []reflect.Type
}

// DatabaseQueryBuilder create an instance of DatabaseQueryBuilder struct.
func NewDatabaseQueryBuilder(types []reflect.Type) *DatabaseQueryBuilder {
	return &DatabaseQueryBuilder{types: types}
}

// CreateAllTable create all table to PostgreSQL database.
func (d *DatabaseQueryBuilder) CreateAllTable() {
	for _, t := range d.types {
		tableName := t.Name()
		fmt.Println(tableName)
		for i := 0; i < t.NumField(); i++ {
			fmt.Println(t.Field(i).Name)
		}
	}
}
