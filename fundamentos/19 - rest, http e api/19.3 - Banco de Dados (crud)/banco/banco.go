package banco

import (
	"database/sql"
	"fmt"

	// import implícito. O pacote mysql vai procurar o driver
	_ "github.com/go-sql-driver/mysql"
)

func Conn() (*sql.DB, error) {
	connectionURL := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True", "root", "root", "db_estudos")
	db, err := sql.Open("mysql", connectionURL)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}

// row, err := db.Query("SELECT * FROM test")

// if err != nil {
// 	log.Fatal(err)
// }

// defer row.Close()

// fmt.Println(row)
