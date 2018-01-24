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
	var Gtimapid int

	query := "SELECT COUNT(*) FROM GTIDOMAIN WHERE GTIDOMAIN.Domain ='" + Domain + "'"
	row := OpenConnection["GTI"].QueryRow(query)
	row.Scan(
		&Count,
	)
	if Count == 0 {

		log.Println("Insert Domain")
		rowss, err := OpenConnection["GTI"].Exec("insert into GTIDOMAIN (Domain,Reputation,Src) values (?,?,?)", Domain, Reputation, source)
		if err != nil {
			log.Println("Error -DB: Profile", err, rowss)
		}

		row1, err := OpenConnection["GTI"].Query("select DomainId from GTIDOMAIN where Domain =?", Domain)
		if err != nil {
			log.Println("Error -DB: User", err)
		}
		for row1.Next() {

			row1.Scan(

				&Gtiid,
			)

		}

		for i := 0; i < len(Categories); i++ {

			row1, err := OpenConnection["GTI"].Query("select ID from GTI where Category =?", Categories[i])
			if err != nil {
				log.Println("Error -DB: User", err)
			}
			for row1.Next() {

				row1.Scan(

					&Gtimapid,
				)

			}

			rows, err := OpenConnection["GTI"].Exec("insert into GTIMAP (DomainID,GTIid) values (?,?)", Gtiid, Gtimapid)
			if err != nil {
				log.Println("Error -DB: Executive insert picture", err, rows)

			}
		}

	}

}
func InsertGti(Domain string, cat int) {
	rowss, err := OpenConnection["GTI"].Exec("insert into GTI (Categoryname,Category) values (?,?)", Domain, cat)
	if err != nil {
		log.Println("Error -DB: Profile", err, rowss)
	}

}
