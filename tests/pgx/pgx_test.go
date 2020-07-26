package pgx

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v4"
    "os"
    "testing"
)

func TestPgx1(t *testing.T) {
    connString :=  "host=172.77.10.15 port=5432 user=test password=6397c7f7f97a dbname=dvdrental"

    var (
        cfg *pgx.ConnConfig
        err error
        conn *pgx.Conn
    )
    cfg, err = pgx.ParseConfig(connString)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
        os.Exit(1)
    }
    cfg.PreferSimpleProtocol = true
    conn, err = pgx.ConnectConfig(context.Background(), cfg)

    if err != nil {
        fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
        os.Exit(1)
    }
    defer conn.Close(context.Background())

    var title string
    var releaseYear int
    err = conn.QueryRow(context.Background(), "select title, release_year from film where film_id=$1", 42).Scan(&title, &releaseYear)
    if err != nil {
        fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("title = %s, year = %d\n", title, releaseYear)
}
