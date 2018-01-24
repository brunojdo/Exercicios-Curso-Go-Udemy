package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:sql@/cursogo")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// update
	smtm, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	smtm.Exec("William", 1)
	smtm.Exec("Ui mudou o nome", 2)

	//delete
	smtm2, _ := db.Prepare("delete from usuarios where id = ?")
	smtm2.Exec(2)

}
