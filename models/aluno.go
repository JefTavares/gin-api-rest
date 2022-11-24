package models

type Aluno struct {
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

// Cria uma lista de Aluno, esta sendo populada lá no main
var Alunos []Aluno
