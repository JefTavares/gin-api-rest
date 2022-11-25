package models

import (
	"fmt"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	RG   string `json:"rg" validate:"len=9, regexp=^[0-9]*$"` // regexp=^[a-zA-Z]*$
	CPF  string `json:"cpf" validate:"len=11, regexp=^[0-9]*$"`
}

// Utiliza o package validator.v2 - Para validar a struct conforme os parametros setado
func ValidaDadoDeAluno(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// Cria uma lista de Aluno, esta sendo populada lá no main
//var Alunos []Aluno
