// AULA 86 - EXECUTANDO INSERTS EM UMA TRANSACAO
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

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuario(id, nome) values (?,?)")

	stmt.Exec(2000, "Bia")    // ALERTA: se estes statments nao forem  validados, e o codigo nao executar rollback
	stmt.Exec(2001, "Carlos") //, entao estas linhas serao executadas

	_, err = stmt.Exec(1, "Tiago") // chave duplicada

	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()
	// quando o commit eh executado, sem o rollback,
	// todos os statements que nao deram erro serao executados

}
