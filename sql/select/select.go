// AULA_88 - EXECUTANDO SELECT E MAPEANDO PARA STRUCT
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func sqlOpen() *sql.DB {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	// fmt.Println(reflect.TypeOf(db))
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	db := sqlOpen()
	defer db.Close()

	rows, _ := db.Query("select id, nome from usuario where id > ?", 5)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
	}

}
