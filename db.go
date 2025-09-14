package main

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

func OpenDB() *sql.DB {
    db, err := sql.Open("sqlite3", "./pharmacy.db")
    if err != nil {
        log.Fatal(err)
    }

    createTables := `
    CREATE TABLE IF NOT EXISTS medicines (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        quantity INTEGER NOT NULL,
        price REAL NOT NULL,
        expiry TEXT NOT NULL,
        supplier TEXT
    );

    CREATE TABLE IF NOT EXISTS sales (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        medicine_id INTEGER,
        quantity INTEGER,
        total REAL,
        date TEXT,
        FOREIGN KEY(medicine_id) REFERENCES medicines(id)
    );

    CREATE TABLE IF NOT EXISTS alerts (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        message TEXT,
        date TEXT
    );
    `
    _, err = db.Exec(createTables)
    if err != nil {
        log.Fatal("Error creating tables:", err)
    }

    return db
}
