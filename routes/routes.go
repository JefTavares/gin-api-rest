package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jeftavares/gin-api-rest/controllers"
)

func HandleRequests() {
	//confitura a aplicação do gin
	r := gin.Default()

	r.GET("/alunos", controllers.ExibeTodosAlunos)

	r.Run() //r.Run(":5000") informa a porta, caso queira
}
