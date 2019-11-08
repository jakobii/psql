package psql

import (
	"bytes"
	"sort"
	"strings"
)

func t(database, schema, table string) string {
	var sql bytes.Buffer
	if database != "" {
		sql.WriteString(`"` + database + `".`)
	}

	if schema != "" {
		sql.WriteString(`"` + schema + `".`)
	} else {
		sql.WriteString(`"dbo".`)
	}
	sql.WriteString(`"` + table + `"`)
	return sql.String()
}

// ToUpdate converts data into a sql update statement.
// values and keys should be fully escaped sql expresions.
func ToUpdate(database, schema, table string, values, keys map[string]string) string {
	var sql bytes.Buffer

	sql.WriteString(`UPDATE ` + t(database, schema, table) + ` SET `)

	var cols = make([]string, 0, len(values))
	for k := range values {
		cols = append(cols, k)
	}
	sort.Strings(cols)

	i := 0
	for _, k := range cols {
		sql.WriteString(`"` + k + `" = ` + values[k])
		if i < (len(cols) - 1) {
			sql.WriteString(`, `)
		}
		i++
	}

	var pks = make([]string, 0, len(keys))
	for k := range keys {
		pks = append(pks, k)
	}
	sort.Strings(pks)

	if len(pks) > 0 {
		sql.WriteString(` WHERE `)
		i = 0
		for _, k := range pks {
			sql.WriteString(`"` + k + `" = ` + keys[k])
			i++
			if i < (len(pks) - 1) {
				sql.WriteString(`AND `)
			}
		}
	}
	sql.WriteString(`;`)
	return sql.String()
}

// ToInsert converts data into a sql insert statement.
func ToInsert(database, schema, table string, values []map[string]string) string {

	var sql bytes.Buffer

	sql.WriteString(`INSERT INTO ` + t(database, schema, table))

	// get all column names
	var cols = make([]string, 0, len(values[0]))
	for _, m := range values {
		for k := range m {
			exists := false
			for _, s := range cols {
				if k == s {
					exists = true
				}
			}
			if !exists {
				cols = append(cols, k)
			}
		}
	}
	sort.Strings(cols)
	sql.WriteString(` ("` + strings.Join(cols, `", "`) + `")`)

	// generate row inserts
	sql.WriteString(` VALUES `)
	i := 0
	for _, m := range values {
		sql.WriteString(`(`)
		j := 0
		for _, k := range cols {
			if v, ok := m[k]; ok {
				sql.WriteString(v)
			} else {
				sql.WriteString(`DEFAULT`)
			}
			if j < (len(cols) - 1) {
				sql.WriteString(`, `)
			}
			j++
		}
		sql.WriteString(`)`)
		if i < (len(values) - 1) {
			sql.WriteString(`, `)
		}
		i++
	}

	sql.WriteString(`;`)
	return sql.String()
}
