package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	url := "root:yourpassword@/initialapirest?charset=utf8&parseTime=True&loc=Local"

	db, e := sql.Open("mysql", url)
	if e != nil {
		return nil, e
	}

	if e = db.Ping(); e != nil {
		return nil, e
	}

	return db, nil
}
