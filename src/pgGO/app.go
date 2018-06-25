package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "os"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "termux"
    password = "123"
    dbname   = "termux"
)

func main() {
    args := os.Args
    fmt.Println(args)
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    sqlStatement := `
        INSERT INTO base (nombre, direccion, telefono, usuario)
        VALUES ($1, $2, $3, $4)
        RETURNING id`
    id := 0
    err = db.QueryRow(sqlStatement, "Richar @daemondev", "SJM", "982929041", "granlinux@gmail.com").Scan(&id)
    if err != nil {
        panic(err)
    }
    fmt.Println("New record ID is:", id)
}
