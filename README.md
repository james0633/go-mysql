# go-mysql

对于 MySQL 数据库常用操作的封闭函数

Requires Go >= 1.10 and MySQL >= 4.1

## Package dbi

### functions

* `insert` insert a new record
* `exe` update of delete
* `row` get a record from result
* `query` get full result from query


### Parameters

* `query` : sql for preparing a statement 
* `args` : sql parameters

### Examples

    var new_id int32 = dbi.insert("INSERT INTO test( b ) VALUES( ? )", 1)


## Notes

### About Time

A zero time.Time argument to Query/Exec is treated as a MySQL zero
timestamp (0000-00-00 00:00:00). A MySQL zero timestamp is returned as
a Go zero time.

Timestamps in MySQL are assumed to be in UTC. time.Time arguments are
stored as UTC and returned as UTC.

### Character Set

Strings are by default UTF-8 encoded in the MySQL connection; they are
automatically converted by the MySQL server as needed.


## Installation

    go get github.com/james0633/go-mysql

## Usage

    import (
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
        _ "github.com/james0633/go-mysql"
    )

    func main() {
        db, err := sql.Open("mysql", "mysql://gotest:gomysql@localhost/test")
        ...
    }

## Testing

    mysql@localhost> grant all on test.* to gotest@localhost;
    mysql@localhost> grant all on test.* to gotest@localhost identified by 'gomysql';

    $ go test
