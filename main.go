package main

import (
	"fmt"
	"github.com/Monteiro712/go-rest-api/db"
	"github.com/Monteiro712/go-rest-api/models"
	"github.com/Monteiro712/go-rest-api/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id:1, Nome:"nome1", Historia:"historia1"},
		{Id:2, Nome:"nome2", Historia:"historia2"},
	}
	
	db.ConectaComBancoDeDados()
	fmt.Println("iniciando")
	routes.HandlerRequest()
}
