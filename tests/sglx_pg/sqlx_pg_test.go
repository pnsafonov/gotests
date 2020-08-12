package sglx_pg

import (
    "context"
    "fmt"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    "log"
    "math/rand"
    "os"
    "testing"
    "time"
)

var (
    ctx = context.Background()
    src rand.Source
    rnd *rand.Rand
)

func init() {
    src = rand.NewSource(time.Now().UnixNano())
    rnd = rand.New(src)
}

type film3 struct {
    title       string
    year        int
    category    string
    lang        string
}

func TestSglx1(t *testing.T) {
    //connString := "host=172.77.10.15 port=5432 user=test password=6397c7f7f97a dbname=dvdrental pool_min_conns=1 pool_max_conns=10"
    connString := "host=172.77.10.15 port=5432 user=test password=6397c7f7f97a dbname=dvdrental sslmode=disable"
    //db, err := sqlx.Connect("pgx", connString)
    db, err := sqlx.Connect("postgres", connString)
    if err != nil {
        log.Fatalln(err)
    }

    q := `
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

    row := db.QueryRow(q, 100)
    f1 := film3{}
    err = row.Scan(&f1.title, &f1.year, &f1.category, &f1.lang)
    if err != nil {
       fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
       os.Exit(1)
    }
    //err = row.Scan(&f1.title, &f1.year, &f1.category, &f1.lang)
    //if err != nil {
    //    fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
    //    os.Exit(1)
    //}
    //err = row.Scan(&f1.title, &f1.year, &f1.category, &f1.lang)
    //if err != nil {
    //   fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
    //   os.Exit(1)
    //}

    rows, err := db.Query(q, 100)
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
