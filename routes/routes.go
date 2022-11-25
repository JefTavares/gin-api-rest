package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jeftavares/gin-api-rest/controllers"
)

func HandleRequests() {
	//confitura a aplicação do gin
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)
	r.Run() //r.Run(":5000") informa a porta, caso queira
}
