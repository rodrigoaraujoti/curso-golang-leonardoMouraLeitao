//
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/cursogo")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("insert into usuario(nome) values(?)")
	stmt.Exec("Maria")
	stmt.Exec("Joao")
	stmt.Exec("Antonio")
	stmt.Exec("Jose")
	stmt.Exec("Paulo")
	stmt.Exec("Ana")

	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId()
	fmt.Println("lastInsertedId", id)

	linhas, _ := res.RowsAffected()
	fmt.Println("rowsAffected", linhas)
}

/* --> erro ao executar o sql.Open com os dados de conexao errados
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x1 addr=0x38 pc=0x4c5867]

goroutine 1 [running]:
sync.(*RWMutex).RLock(...)
	C:/Go/src/sync/rwmutex.go:48
database/sql.(*Stmt).ExecContext(0x0, 0x684da0, 0xc000064080, 0xc00008df78, 0x1, 0x1, 0x0, 0x0, 0x0, 0x0)
	C:/Go/src/database/sql/sql.go:2307 +0x57
database/sql.(*Stmt).Exec(...)
	C:/Go/src/database/sql/sql.go:2336
main.main()
	c:/projetos/curso-golang-leonardoMouraLeitao/sql/insert/insert.go:19 +0x158
exit status 2
*/
