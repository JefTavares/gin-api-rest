package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jeftavares/gin-api-rest/models"
)

/*Certo! O não recebimento do contexto do Gin por parâmetro faz com que não seja possível retornar o JSON. O contexto é muito importante e possui informações úteis, como por exemplo, cabeçalhos, cookies, etc.*/
func ExibeTodosAlunos(c *gin.Context) {
	//c.JSON(200, gin.H{		"id":   "1",		"nome": "Jef Tavares",	})
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "E ai " + nome + ", tudo beleza?",
	})
}
