package pgx

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v4"
    "github.com/jackc/pgx/v4/pgxpool"
    "os"
    "testing"
    "time"
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

    count := 100

    for i := 0; i < count; i++ {
        var title string
        var releaseYear int
        err = conn.QueryRow(context.Background(), "select title, release_year from film where film_id=$1", 42).Scan(&title, &releaseYear)
        if err != nil {
            fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
            //os.Exit(1)
            time.Sleep(time.Second * 5)
        }

        fmt.Printf("title = %s, year = %d\n", title, releaseYear)
    }


}

func TestPgx2(t *testing.T) {
    connString :=  "host=172.77.10.15 port=5432 user=test password=6397c7f7f97a dbname=dvdrental pool_min_conns=1 pool_max_conns=10"
    var (
        poolCfg *pgxpool.Config
        pool    *pgxpool.Pool
        err error
    )

    poolCfg, err = pgxpool.ParseConfig(connString)
    if err != nil {
        fmt.Fprintf(os.Stderr, "err: %v\n", err)
        os.Exit(1)
    }

    //poolCfg.ConnConfig.PreferSimpleProtocol = true

    poolCfg, err = pgxpool.ParseConfig(connString)
    if err != nil {
        fmt.Fprintf(os.Stderr, "err: %v\n", err)
        os.Exit(1)
    }

    ctx := context.Background()
    pool, err = pgxpool.ConnectConfig(ctx, poolCfg)
    if err != nil {
        fmt.Fprintf(os.Stderr, "err: %v\n", err)
        os.Exit(1)
    }

    count := 100

    for i := 0; i < count; i++ {
        var title string
        var releaseYear int
        err = pool.QueryRow(context.Background(), "select title, release_year from film where film_id=$1", 42).Scan(&title, &releaseYear)
        if err != nil {
            fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
            //os.Exit(1)
            time.Sleep(time.Second * 1)
            continue
        }

        fmt.Printf("title = %s, year = %d\n", title, releaseYear)
    }


}