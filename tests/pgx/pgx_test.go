package pgx

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v4"
    "github.com/jackc/pgx/v4/pgxpool"
    "log"
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

func getConnPool0(connString string) *pgxpool.Pool {
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
    return pool
}

func getConnPool1() *pgxpool.Pool {
    connString := "host=172.77.10.15 port=5432 user=test password=6397c7f7f97a dbname=dvdrental pool_min_conns=1 pool_max_conns=10"
    return getConnPool0(connString)
}

func getConnPool2() *pgxpool.Pool {
    connString := "host=172.77.10.15 port=5432 user=test password=6397c7f7f97a dbname=test pool_min_conns=1 pool_max_conns=10"
    return getConnPool0(connString)
}

func TestPgx2(t *testing.T) {
    pool := getConnPool1()

    ctx := context.Background()
    count := 100
    for i := 0; i < count; i++ {
        var title string
        var releaseYear int
        err := pool.QueryRow(ctx, "select title, release_year from film where film_id=$1", 42).Scan(&title, &releaseYear)
        if err != nil {
            fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
            //os.Exit(1)
            time.Sleep(time.Second * 1)
            continue
        }

        fmt.Printf("title = %s, year = %d\n", title, releaseYear)
    }
}

type film3 struct {
    title       string
    year        int
    category    string
    lang        string
}

func (f *film3) load(rows pgx.Rows) (bool, error) {
    if !rows.Next() {
        return false, nil
    }
    err := rows.Scan(&f.title, &f.year, &f.category, &f.lang)
    if err != nil {
        return false, err
    }
    return true, nil
}


func TestPgx3(t *testing.T) {
    pool := getConnPool1()

    q :=
`
SELECT t0.title, t0.release_year, t2.name, t3.name
FROM film as t0
JOIN film_category as t1
	ON t0.film_id = t1.film_id
JOIN category as t2
	ON t1.category_id = t2.category_id
JOIN language as t3
	ON t0.language_id = t3.language_id
ORDER BY t0.film_id desc
LIMIT $1
`
    ctx := context.Background()
    rows, err := pool.Query(ctx, q, 100)
    if err != nil {
        fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
        os.Exit(1)
    }

    films := make([]film3, 0, 16)

    for rows.Next() {
        f := film3{}
        err := rows.Scan(&f.title, &f.year, &f.category, &f.lang)
        if err != nil {
            fmt.Fprintf(os.Stderr, "scan failed: %v\n", err)
            os.Exit(1)
        }
        films = append(films, f)
    }

    log.Println("done")
}

func TestContext1(t *testing.T) {
    contextMain := context.Background()

    //c1, f1 := context.WithTimeout(contextMain, time.Second * 1000)
    //f1()

    //c1, _ := context.WithTimeout(contextMain, time.Second * 1000)

    c1, _ := context.WithTimeout(contextMain, time.Second * 1)
    time.Sleep(time.Second * 2)

    сh1 := c1.Done()
    _, ok := <- сh1
    fmt.Printf("ok = %v\n", ok)
    _, ok = <- сh1
    fmt.Printf("ok = %v\n", ok)
    _, ok = <- сh1
    fmt.Printf("ok = %v\n", ok)
}

func TestContext2(t *testing.T) {
    c0 := context.Background()

    c1, _ := context.WithTimeout(c0, time.Second * 1)
    time.Sleep(time.Second * 2)

    ch0 := c0.Done()
    ch1 := c1.Done()

    _, ok1 := <- ch1
    fmt.Printf("ok = %v\n", ok1)
    _, ok0 := <- ch0
    fmt.Printf("ok = %v\n", ok0)
}

func TestContext3(t *testing.T) {
    c0 := context.Background()

    c1, _ := context.WithTimeout(c0, time.Second * 1000)
    c2, _ := context.WithTimeout(c0, time.Second * 1)

    ch0 := c0.Done()
    ch1 := c1.Done()
    ch2 := c2.Done()

    time.Sleep(time.Second * 2)

    _, ok2 := <- ch2
    fmt.Printf("ok = %v\n", ok2)
    _, ok1 := <- ch1
    fmt.Printf("ok = %v\n", ok1)
    _, ok0 := <- ch0
    fmt.Printf("ok = %v\n", ok0)
}

func TestContext4(t *testing.T) {
    c0 := context.Background()

    c1, _ := context.WithTimeout(c0, time.Second * 5)
    c2, _ := context.WithTimeout(c0, time.Second * 1)

    //ch0 := c0.Done()
    ch1 := c1.Done()
    ch2 := c2.Done()

    time.Sleep(time.Second * 2)

    _, ok2 := <- ch2
    fmt.Printf("ok = %v\n", ok2)
    _, ok1 := <- ch1
    fmt.Printf("ok = %v\n", ok1)
    //_, ok0 := <- ch0
    //fmt.Printf("ok = %v\n", ok0)
}

func TestPgx4(t *testing.T) {
    //pool := getConnPool2()

    //pool.Begin()
}