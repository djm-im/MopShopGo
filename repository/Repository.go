package repository

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

type Repository struct {
    database *sql.DB
}

var mysqlDatabase = Repository{
    database: connectDb(),
}

func connectDb() *sql.DB {
    database, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/schema_MopShop")

    if err != nil {
        log.Fatalf("Cannot conntect to database. Terminated.")
    }

    return database
}

func GetMysql() *sql.DB {
    return mysqlDatabase.database
}
