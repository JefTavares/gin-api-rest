package main

import "github.com/gin-gonic/gin"

/*Certo! O não recebimento do contexto do Gin por parâmetro faz com que não seja possível retornar o JSON. O contexto é muito importante e possui informações úteis, como por exemplo, cabeçalhos, cookies, etc.*/
func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "Jef Tavares",
	})
}

func main() {
	//confitura a aplicação do gin
	r := gin.Default()

	r.GET("/alunos", ExibeTodosAlunos)

	r.Run() //r.Run(":5000") informa a porta, caso queira
}
