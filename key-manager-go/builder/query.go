package builder

import (
	"reflect"
	"strings"

	"github.com/ddiogoo/broker/tree/master/key-manager-go/util"
)

var (
	// h storage all database types.
	h = make(map[string]string)
)

type DatabaseQueryBuilder struct {
	types  []reflect.Type
	Tables []string
}

// DatabaseQueryBuilder create an instance of DatabaseQueryBuilder struct.
func NewDatabaseQueryBuilder(types []reflect.Type) *DatabaseQueryBuilder {
	return &DatabaseQueryBuilder{types: types}
}

// BuildCreateTable create all table query to PostgreSQL database.
func (d *DatabaseQueryBuilder) BuildCreateTable() {
	d.storageDatabaseType()
	for _, t := range d.types {
		var sb strings.Builder
		sb.WriteString("CREATE TABLE " + t.Name() + " ( ")
		for i := 0; i < t.NumField(); i++ {
			prop := t.Field(i).Name
			value := t.Field(i).Type.Name()
			sb.WriteString(prop + " " + h[value] + ", ")
		}
		sb.WriteString(");")
		result := util.ReplaceLastOccurrence(sb.String(), ",", "")
		d.Tables = append(d.Tables, result)
	}
}

// storageDatabaseType charge all database types.
func (d *DatabaseQueryBuilder) storageDatabaseType() {
	h["string"] = "VARCHAR(255)"
}
