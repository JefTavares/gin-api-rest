package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeftavares/gin-api-rest/database"
	"github.com/jeftavares/gin-api-rest/models"
)

/*Certo! O não recebimento do contexto do Gin por parâmetro faz com que não seja possível retornar o JSON. O contexto é muito importante e possui informações úteis, como por exemplo, cabeçalhos, cookies, etc.*/
func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, &alunos)
	//c.JSON(200, gin.H{		"id":   "1",		"nome": "Jef Tavares",	})
	//c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "E ai " + nome + ", tudo beleza?",
	})
}

func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidaDadoDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Aluno " + id + " não encontrado"})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func DeleteAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.ValidaDadoDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)

	c.JSON(http.StatusOK, aluno)

}

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno) //pesquisa e joga na variavel aluno

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"Not found": "Aluno " + cpf + " não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}
