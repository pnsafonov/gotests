package pgx

import (
    "context"
    "fmt"
    "github.com/jackc/pgx/v4"
    "github.com/jackc/pgx/v4/pgxpool"
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
    pool := getConnPool2()

    q :=
`
DROP TABLE IF EXISTS tuser;

CREATE TABLE tuser (
	id 				SERIAL PRIMARY KEY,
	login 			TEXT UNIQUE NOT NULL CHECK (login <> ''),
	email 			TEXT UNIQUE NOT NULL CHECK (email <> ''),
	mobile_phone	TEXT UNIQUE
);
`

    ct, err := pool.Exec(ctx, q)
    if err != nil {
        t.Fatalf("err = %v", err)
    }

    fmt.Printf("ct = %s\n", ct)
    fmt.Printf("done")
}

func TestPgx5(t *testing.T) {
    pool := getConnPool2()

    q :=
`
    INSERT INTO tuser
    (login, email, mobile_phone)
    VALUES ($1, $2, $3)
    RETURNING id
`

    count := 10
    for i := 0; i < count; i++ {
        n := rnd.Int31()
        login := fmt.Sprintf("login_%d", n)
        email := fmt.Sprintf("name_%d@gmail.com", n)
        phone := fmt.Sprintf("+7-950-133-74-80-%d", n)

        row := pool.QueryRow(ctx, q, login, email, phone)

        var id int
        err := row.Scan(&id)
        if err != nil {
            t.Fatalf("err = %v", err)
        }
    }

    fmt.Println("done")
}

type sys6 struct {
    id          int
    dbVersion   int
    appVersion  string
}

// cannot insert multiple commands into a prepared statement - QueryRow
// -> poolCfg.ConnConfig.PreferSimpleProtocol = true
// no rows in results set - Scan
func TestPgx6(t *testing.T) {
    pool := getConnPool2()

    q :=
`
BEGIN TRANSACTION;
    
CREATE TABLE IF NOT EXISTS sys (
	id 				SERIAL PRIMARY KEY,
    db_version      INTEGER NOT NULL,
	app_version 	TEXT UNIQUE NOT NULL CHECK (app_version <> '')
);

INSERT INTO sys
(id, db_version, app_version)
VALUES (1, $1, $2)
ON CONFLICT (id) DO UPDATE
	SET 
	db_version = $1, 
	app_version = $2
RETURNING id, db_version, app_version;

COMMIT TRANSACTION;
`

    dbVersion := 3
    appVersion := "1.0.3"
    row := pool.QueryRow(ctx, q, dbVersion, appVersion)

    sys := sys6{}
    err := row.Scan(&sys.id, &sys.dbVersion, &sys.appVersion)
    if err != nil {
        t.Fatalf("err = %v", err)
    }

    fmt.Printf("id = %d, db_version = %d, app_version = %s\n", sys.id, sys.dbVersion, sys.appVersion)
    fmt.Println("done")
}

// cannot insert multiple commands into a prepared statement
func TestPgx7(t *testing.T) {
    pool := getConnPool2()

    q :=
` 
CREATE TABLE IF NOT EXISTS sys (
	id 				SERIAL PRIMARY KEY,
    db_version      INTEGER NOT NULL,
	app_version 	TEXT UNIQUE NOT NULL CHECK (app_version <> '')
);

INSERT INTO sys
(id, db_version, app_version)
VALUES (1, $1, $2)
ON CONFLICT (id) DO UPDATE
	SET 
	db_version = $1, 
	app_version = $2
RETURNING id, db_version, app_version;
`

    tx, err := pool.Begin(ctx)
    if err != nil {
        t.Fatalf("err = %v", err)
    }

    dbVersion := 3
    appVersion := "1.0.3"
    row := tx.QueryRow(ctx, q, dbVersion, appVersion)

    sys := sys6{}
    err = row.Scan(&sys.id, &sys.dbVersion, &sys.appVersion)
    if err != nil {
        t.Fatalf("err = %v", err)
    }

    err = tx.Commit(ctx)
    if err != nil {
        t.Fatalf("err = %v", err)
    }

    fmt.Printf("id = %d, db_version = %d, app_version = %s\n", sys.id, sys.dbVersion, sys.appVersion)
    fmt.Println("done")
}

func TestPgx8(t *testing.T) {
    pool := getConnPool2()

    tx, err := pool.Begin(ctx)
    if err != nil {
        t.Fatalf("err = %v", err)
    }

    q0 :=
`
CREATE TABLE IF NOT EXISTS sys (
	id 				SERIAL PRIMARY KEY,
    db_version      INTEGER NOT NULL,
	app_version 	TEXT UNIQUE NOT NULL CHECK (app_version <> '')
);
`

    // no rows in result set
    //row0 := tx.QueryRow(ctx, q0)
    //err = row0.Scan()
    //if err != nil {
    //    t.Fatalf("err = %v", err)
    //}

    _, err = tx.Exec(ctx, q0)
    if err != nil {
       t.Fatalf("err = %v", err)
    }
    //fmt.Printf("ct = %s\n", ct)
    //ct.RowsAffected()

    q :=
`
INSERT INTO sys
(id, db_version, app_version)
VALUES (1, $1, $2)
ON CONFLICT (id) DO UPDATE
	SET 
	db_version = $1, 
	app_version = $2
RETURNING id, db_version, app_version;
`

    dbVersion := 4
    appVersion := "1.0.4"
    row := tx.QueryRow(ctx, q, dbVersion, appVersion)

    sys := sys6{}
    err = row.Scan(&sys.id, &sys.dbVersion, &sys.appVersion)
    if err != nil {
        t.Fatalf("err = %v", err)
    }

    err = tx.Commit(ctx)
    if err != nil {
        t.Fatalf("err = %v", err)
    }

    fmt.Printf("id = %d, db_version = %d, app_version = %s\n", sys.id, sys.dbVersion, sys.appVersion)
    fmt.Println("done")
}

func TestPgx9(t *testing.T) {
    pool := getConnPool2()

    q :=
`
SELECT 'public.sys'::regclass
`
    row := pool.QueryRow(ctx, q)
    var name1 string
    err := row.Scan(&name1)
    if err != nil {
        t.Fatalf("err = %v", err)
    }

//    q2 :=
//`
//SELECT 'public.'$1::regclass
//`
//    row2 := pool.QueryRow(ctx, q2, "sys1")
//    var name2 string
//    err = row2.Scan(name2)
//    if err != nil {
//        t.Fatalf("err = %v", err)
//    }

    q2 :=
        `
SELECT 'public.sys111'::regclass
`
    row2 := pool.QueryRow(ctx, q2)
    var name2 string
    err = row2.Scan(name2)
    if err != nil {
       t.Fatalf("err = %v", err)
    }

    log.Println("done")
}