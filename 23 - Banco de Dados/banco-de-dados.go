package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	conexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, error := sql.Open("mysql", conexao)
	if error != nil {
		fmt.Println("Erro ao chamar o Open de conexão!!")
		log.Fatal(error)
	}
	defer db.Close()

	if error := db.Ping(); error != nil {
		fmt.Print("Erro ao chamar o Ping na conexão!!")
		log.Fatal(error)
	}

	fmt.Println("Conexão aberta!")

	linhas, error := db.Query("SELECT * FROM usuarios")
	if error != nil {
		fmt.Print("Houve um erro ao executar o SELECT SQL")
		log.Fatal(error)
	}
	defer linhas.Close()

	fmt.Println(linhas)

}
