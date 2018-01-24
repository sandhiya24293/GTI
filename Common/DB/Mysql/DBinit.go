package Mysql

import (
	InitDb "Dgraph/Common/DB/Mysql/InitializeDB"
	"database/sql"
)

var OpenConnection = make(map[string]*sql.DB)

func init() {
	OpenConnection = InitDb.Ret()
}
