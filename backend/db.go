package main

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
    var err error
    // Use your MySQL username, password, host, port, and database name
    DB, err = sql.Open("mysql", "lungile:1234@tcp(127.0.0.1:3306)/go_dev")
    if err != nil {
        log.Fatal("DB open error:", err)
    }

    // Check connection
    if err = DB.Ping(); err != nil {
        log.Fatal("DB ping error:", err)
    }

    log.Println("Connected to MySQL database successfully!")
}


