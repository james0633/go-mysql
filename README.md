# go-mysql

对于 MySQL 数据库常用操作的封闭函数

Requires Go >= 1.10 and MySQL >= 4.1

## Package dbi

### functions

* `add` insert a new record
* `exe` update of delete
* `row` get one record from result
* `all` get all result from query


### Parameters

* `query` : sql for preparing a statement 
* `args` : sql parameters

### Examples

    dbo.MyDB, err = sql.Open("mysql", config.MysqlDSN)
    if err != nil {
        panic(err.Error())
    }
    //defer dbo.MyDB.Close()
    dbo.MyDB.SetMaxOpenConns(config.MysqlMaxConn)
    dbo.MyDB.SetMaxIdleConns(config.MysqlMaxIdle)
    dbo.MyDB.Ping()
    
    var new_id int32 = dbo.add("INSERT INTO test( b ) VALUES( ? )", 1)
    
    row, err := dbo.Row("SELECT * FROM testA WHERE id=?", 2019)
    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("find record: ", row["id"])
    }


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
        "github.com/james0633/go-mysql"
    )

    func main() {
        dbo.MyDB, err := sql.Open("mysql", "mysql://gotest:gomysql@localhost/test")
        ...
    }

## Testing

    mysql@localhost> grant all on test.* to gotest@localhost;
    mysql@localhost> grant all on test.* to gotest@localhost identified by 'gomysql';

    $ go test
