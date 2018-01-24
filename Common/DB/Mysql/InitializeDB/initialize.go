package InitializeDB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var OpenConnection = make(map[string]*sql.DB)

func init() {
	Tracker, err := sql.Open("mysql", "root:mypassword@tcp(172.17.0.6:3306)/GTIRC?charset=utf8")

	if err != nil {
		fmt.Println("error", err)
	}

	OpenConnection["GTI"] = Tracker

}
func Ret() map[string]*sql.DB {
	return OpenConnection
}
