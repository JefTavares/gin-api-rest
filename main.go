package main

import (
	"github.com/jeftavares/gin-api-rest/models"
	"github.com/jeftavares/gin-api-rest/routes"
)

func main() {

	//a var Alunos do tipo array da struct Aluno
	models.Alunos = []models.Aluno{
		{Nome: "Jef Tavares", CPF: "327.292.678-08", RG: "42.355.815-8"},
		{Nome: "Thete Gatuzo", CPF: "123.321.876-11", RG: "12.234.322-1"},
		{Nome: "Ju Gatuzo", CPF: "111.222.333-44", RG: "11.222.333-4"},
	}

	routes.HandleRequests()
}
