package sqlx_mysql

import (
    "context"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "log"
    "math/rand"
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

func logFatal(err error) {
    log.Fatalf("err = %v\n", err)
}

func getDb0() *sqlx.DB {
    connString := "test:qfiBek3vhgo5@tcp(172.77.20.10:3306)/test?parseTime=true&charset=utf8"
    db, err := sqlx.Connect("mysql", connString)
    if err != nil {
        logFatal(err)
    }
    return db
}

type AddPhoneRequest struct {
    AgentID 		   int32  `db:"agent_id"`
    Number 			   string `db:"number"`
    Name 			   string `db:"name"`
    BlockInternational bool	  `db:"block_international"`
}

func TestSqlxMysql1(t *testing.T) {
    db := getDb0()

    q := `
INSERT INTO cc_agent_number 
(id_agent, number, name, block_international)
VALUES (:agent_id, :number, :name, :block_international)
ON DUPLICATE KEY UPDATE
	id_agent=:agent_id, name=:name, block_international=:block_international;
`

    n := rnd.Int31()
    req := AddPhoneRequest{}
    req.AgentID = 1
    req.Number = fmt.Sprintf("8-930-560-12-55-%d", n)
    req.Name = fmt.Sprintf("name_%d", n)
    req.BlockInternational = (n % 2) == 0

    // sql: converting argument $1 type: unsupported type sqlx_mysql.AddPhoneRequest, a struct
    //result, err := db.Exec(q, req)

    result, err := db.NamedExec(q, req)
    if err != nil {
        logFatal(err)
    }

    ra, err := result.RowsAffected()
    if err != nil {
        logFatal(err)
    }
    log.Printf("ra = %d\n", ra)

    log.Println("done")
}
