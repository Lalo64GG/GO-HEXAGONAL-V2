package database

import (
    "database/sql"
    "log"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    var err error
    
    dsn := os.Getenv("DB_CONN")
    if dsn == "" {
        log.Fatalf("DB_CONN environment variable not set")
    }

    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error opening database connection: %v", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatalf("Error pinging database: %v", err)
    }

    log.Println("Database connection established successfully")
}

func WithTransaction(db *sql.DB, fn func(tx *sql.Tx) error) error {
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    defer func() {
        if p := recover(); p != nil {
            tx.Rollback()
            panic(p)
        } else if err != nil {
            tx.Rollback()
        } else {
            err = tx.Commit()
        }
    }()

    err = fn(tx)
    return err
}
