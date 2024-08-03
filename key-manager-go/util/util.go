package util

import (
	"reflect"
	"strings"

	"github.com/ddiogoo/broker/tree/master/key-manager-go/model"
)

// ReplaceLastOccurrence change the last ocurrence to another string.
func ReplaceLastOccurrence(s, old, new string) string {
	pos := strings.LastIndex(s, old)
	if pos == -1 {
		return s
	}
	return s[:pos] + new + s[pos+len(old):]
}

// BuildListOfReflectType create a list with all models.
func BuildListOfReflectType() []reflect.Type {
	typs := []reflect.Type{}
	typs = append(typs, reflect.TypeOf(&model.Key{}).Elem())
	return typs
}
