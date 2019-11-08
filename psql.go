package psql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // comment
)

// Connection contains the variables required to create a sql server connection.
type Connection struct {
	Server   string
	Database string
	Username string
	Password string
	Port     int
}

// NewConnection creates a new Connection with some default values.
func NewConnection() Connection {
	var c Connection
	c.Port = 1433
	return c
}

func (c Connection) String() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.Username, c.Password, c.Server, c.Database)
}

// Column represents a sql tables column configuration.
type Column struct {
	Name         string
	Position     int
	Type         string
	Nullable     bool
	Key          bool
	Unique       bool
	Default      string
	CharacterSet string
	Percision    int
	Scale        int
	Length       int
}

// GetColumns gets a tables column configuration.
func GetColumns(c Connection, schema, table string) ([]Column, error) {

	// open db
	db, err := sql.Open("postgres", c.String())
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// generate sql
	query := fmt.Sprintf(
		sqlGetColumns,
		Escape(schema),
		Escape(table),
	)

	// exec sql
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// build return object
	var columns = make([]Column, 1)
	for rows.Next() {
		var col Column
		err := rows.Scan(
			&col.Name,
			&col.Position,
			&col.Type,
			&col.Nullable,
			&col.Key,
			&col.Unique,
			&col.Default,
			&col.CharacterSet,
			&col.Percision,
			&col.Scale,
			&col.Length,
		)
		if err != nil {
			return nil, err
		}
		columns = append(columns, col)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return columns, nil
}

//// Query performs a tsql query and returns the results if any.
//func Query(conn Connection, query string) ([]map[string]interface{}, error) {
//
//	// get table schema
//
//}
//
//// Get retrieves a single row by primary key. if columns is nill all columns will be returned. if the resulting query returns more then one row an error is thrown.
//func Get(conn Connection, schema, table string, columns []string, keys map[string]interface{}) (map[string]interface{}, error) {
//
//}
//
//// Set inserts an updates a single row. if the tables primary keys are included in the values parameter they will be updated.
//func Set(conn Connection, schema, table string, values map[string]interface{}, keys map[string]interface{}) error {
//
//}
//
//// Del deletes a single row. the primary keys must be included in the values parameter.
//func Del(conn Connection, schema, table string, keys map[string]interface{}) error {
//
//}
//}
//
//// Del deletes a single row. the primary keys must be included in the values parameter.
//func Del(conn Connection, schema, table string, keys map[string]interface{}) error {
//
//}
