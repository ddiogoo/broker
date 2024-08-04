package builder

import (
	"fmt"
	"reflect"
	"strings"
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
		sb.WriteString("CREATE TABLE " + t.Name() + " (")
		for i := 0; i < t.NumField(); i++ {
			prop := t.Field(i).Name
			value := t.Field(i).Type.Name()
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(prop + " " + h[value])
		}
		sb.WriteString(");")
		d.Tables = append(d.Tables, sb.String())
	}
}

// BuildInsert creates a insert query to execute on database.
func (d *DatabaseQueryBuilder) BuildInsert(obj interface{}) string {
	rt := reflect.TypeOf(obj)
	rv := reflect.ValueOf(obj)
	var sb strings.Builder

	sb.WriteString("INSERT INTO " + rt.Name() + " (")
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if !field.IsExported() {
			continue
		}
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(field.Name)
	}

	sb.WriteString(") VALUES (")
	for i := 0; i < rv.NumField(); i++ {
		field := rt.Field(i)
		if !field.IsExported() {
			continue
		}
		fieldValue := rv.Field(i).Interface()
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("'%v'", fieldValue))
	}

	sb.WriteString(");")
	return sb.String()
}

// storageDatabaseType charge all database types.
func (d *DatabaseQueryBuilder) storageDatabaseType() {
	h["string"] = "VARCHAR(255)"
}
