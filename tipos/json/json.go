// AULA_062 - JSON
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promocao", "Eletronico"}}
	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))
	// fmt.Println(p1Json) // slice de byte

	// json para struct
	var p2 produto
	// jsonString := `{"id":2,"nome":"Caneta","preco":8.90,"tags":["Papelaria","Importado"]}`
	jsonString := `{"id":2,"nome":"Caneta","preco":8.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags)
	fmt.Println(reflect.TypeOf(jsonString)) // string?

}
