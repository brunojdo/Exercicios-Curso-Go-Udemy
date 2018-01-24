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

	tx, _ := db.Begin()

	stmt, _ := tx.Prepare("insert into usuarios(id,nome) values(?,?)")

	stmt.Exec(1001, "Ana")
	stmt.Exec(2002, "Zezé")
	_, err = stmt.Exec(10, "Esse não") //chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal("Ops")
	}

	tx.Commit()

}
