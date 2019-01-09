package dbo

import (
    "strings"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var MyDB *sql.DB

//add new record to table
func Add(query string, args ...interface{}) (int64, error) {
    stmt, err := MyDB.Prepare(query)
    if err != nil {
        panic(err.Error())
    }
    defer stmt.Close()

    result, err := stmt.Exec(args...)
    if err != nil {
        panic(err.Error())
    }
    return result.LastInsertId()
}

//execute commands
func Exe(query string, args ...interface{}) (int64, error) {
    stmt, err := MyDB.Prepare(query)
    if err != nil {
        panic(err.Error())
    }
    defer stmt.Close()

    result, err := stmt.Exec(args...)
    if err != nil {
        panic(err.Error())
    }
    return result.RowsAffected()
}

//get one record from query result
func Row(query string, args ...interface{}) (map[string]string, error) {
    if !strings.Contains(strings.ToUpper(query), "LIMIT") {
        query += " LIMIT 1"
    }
    stmt, err := MyDB.Prepare(query)
    if err != nil {
        panic(err.Error())
    }
    defer stmt.Close()

    rows, err := stmt.Query(args...)
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()

    columns, err := rows.Columns()
    if err != nil {
        panic(err.Error())
    }

    values := make([]sql.RawBytes, len(columns))
    scanArgs := make([]interface{}, len(values))
    ret := make(map[string]string, len(scanArgs))

    for i := range values {
        scanArgs[i] = &values[i]
    }
    for rows.Next() {
        err = rows.Scan(scanArgs...)
        if err != nil {
            panic(err.Error())
        }
        var value string

        for i, col := range values {
            if col == nil {
                value = "" //or NULL
            } else {
                value = string(col)
            }
            ret[columns[i]] = value
        }
        break //get the first row only
    }
    return ret, nil
}

//get all records from query result
func All(query string, args ...interface{}) ([]map[string]string, error) {
    stmt, err := MyDB.Prepare(query)
    if err != nil {
        panic(err.Error())
    }
    defer stmt.Close()

    rows, err := stmt.Query(args...)
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()

    columns, err := rows.Columns()
    if err != nil {
        panic(err.Error())
    }

    values := make([]sql.RawBytes, len(columns))
    scanArgs := make([]interface{}, len(values))

    ret := make([]map[string]string, 0)
    for i := range values {
        scanArgs[i] = &values[i]
    }

    for rows.Next() {
        err = rows.Scan(scanArgs...)
        if err != nil {
            panic(err.Error())
        }
        var value string
        vmap := make(map[string]string, len(scanArgs))
        for i, col := range values {
            if col == nil {
                value = "" // or NULL
            } else {
                value = string(col)
            }
            vmap[columns[i]] = value
        }
        ret = append(ret, vmap)
    }
    return ret, nil
}
