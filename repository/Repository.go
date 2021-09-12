package repository

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

const (
    dbUser = "root"
    dbPass = "password"
    host   = "127.0.0.1"
    port   = "3306"
    schema = "schema_MopShop"
)

type repository struct {
    database *sql.DB
}

var mysqlDatabase = repository{
    database: connectDb(),
}

func connectDb() *sql.DB {
    database, err := sql.Open(
        "mysql",
        fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, host, port, schema))

    if err != nil {
        log.Fatalf("Cannot conntect to database. Terminated.")
    }

    return database
}

func GetMysql() *sql.DB {
    return mysqlDatabase.database
}
