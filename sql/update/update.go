// AULA_86 - EXECUTANDO UPDATE E DELETE
package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//update
	stmt, _ := db.Prepare("update usuario set nome = ? where id = ?")
	stmt.Exec("Uóxiton Istive", 1)
	stmt.Exec("Sheristone Uasleska", 2)

	//delete
	stmt2, _ := db.Prepare("delete from usuario where id = ?")
	stmt2.Exec(3)
}