package dbi

import (
    "strings"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//insert new record to table
func insert(query string, args ...interface{}) (int64, error) {
    stmt, err := db.Prepare(query)
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
func exe(query string, args ...interface{}) (int64, error) {
    stmt, err := db.Prepare(query)
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
func row(query string, args ...interface{}) (*map[string]string, error) {
    if !strings.Contains(strings.ToUpper(query), "LIMIT") {
        query += " LIMIT 1"
    }
    stmt, err := db.Prepare(query)
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
    return &ret, nil
}

//get all records from query result
func query(query string, args ...interface{}) (*[]map[string]string, error) {
    stmt, err := db.Prepare(query)
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
    return &ret, nil
}

