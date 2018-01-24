package Mysql

import (
	_ "database/sql"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
	Categories []int
}

func InsertDomain(Domain string, Reputation int, Categories []int, source string) {
	log.Println("inside")
	var Count int
	var Gtiid int

	query := "SELECT COUNT(*) FROM GTIDOMAIN WHERE GTIDOMAIN.Domain_name ='" + Domain + "'"
	row := OpenConnection["GTI"].QueryRow(query)
	row.Scan(
		&Count,
	)
	if Count == 0 {

		log.Println("Insert Domain")
		rowss, err := OpenConnection["GTI"].Exec("insert into GTIDOMAIN (Domain_name,Reputation,Sourcefile) values (?,?,?)", Domain, Reputation, source)
		if err != nil {
			log.Println("Error -DB: Profile", err, rowss)
		}

		row1, err := OpenConnection["GTI"].Query("select Domain_ID from GTIDOMAIN where Domain_name =?", Domain)
		if err != nil {
			log.Println("Error -DB: User", err)
		}
		for row1.Next() {

			row1.Scan(

				&Gtiid,
			)

		}

		for i := 0; i < len(Categories); i++ {

			rows, err := OpenConnection["GTI"].Exec("insert into GTIDOMAINCATMAPPING (Domain_ID,Category) values (?,?)", Gtiid, Categories[i])
			if err != nil {
				log.Println("Error -DB: Executive insert picture", err, rows)

			}
		}

	}

}
func InsertGti(Domain string, cat int) {
	rowss, err := OpenConnection["GTI"].Exec("insert into GTI (Category,Category_name) values (?,?)", cat, Domain)
	if err != nil {
		log.Println("Error -DB: Profile", err, rowss)
	}

}
