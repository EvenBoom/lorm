package lorm

import (
	"reflect"
	"strconv"
	"strings"
)

//Query is a method for query a struct from db
func Query(s interface{}) error {
	elem := reflect.ValueOf(s).Elem()
	table := strings.Split(elem.Type().String(), ".")
	//This sql sentence is just for mysql
	rows, err := db.Query("select * from " + table[len(table)-1] + " limit 0," + strconv.Itoa(elem.Len()))

	if elem.Kind() == reflect.Slice {
		n := elem.Index(0).NumField()
		for i := 0; rows.Next(); i++ {
			index := elem.Index(i)
			v := make([]reflect.Value, n)
			for j := 0; j < n; j++ {
				f := index.Field(j)
				v[j] = f.Addr()
			}
			meth := reflect.ValueOf(rows.Scan)
			meth.Call(v)
		}

	}
	return err
}

//Query is a method for query a struct from db with condition
func QueryByCondition(s interface{},condition string,args ...interface{}) error {
	elem := reflect.ValueOf(s).Elem()
	table := strings.Split(elem.Type().String(), ".")
	//This sql sentence is just for mysql
	rows, err := db.Query("select * from " + table[len(table)-1] + " " + condition + " limit 0," + strconv.Itoa(elem.Len()),args ...)

	if elem.Kind() == reflect.Slice {
		n := elem.Index(0).NumField()
		for i := 0; rows.Next(); i++ {
			index := elem.Index(i)
			v := make([]reflect.Value, n)
			for j := 0; j < n; j++ {
				f := index.Field(j)
				v[j] = f.Addr()
			}
			meth := reflect.ValueOf(rows.Scan)
			meth.Call(v)
		}

	}
	return err
}
