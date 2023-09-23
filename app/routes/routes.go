package routes

import (
	"github.com/gin-gonic/gin"
	"golang-first-api-rest/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.POST("/pessoas", controllers.CriaPessoa)
	r.GET("/pessoas/:id", controllers.BuscaPessoaPorId)
	r.GET("/pessoas", controllers.BuscaPessoasPorTermo)
	r.GET("/contagem-pessoas", controllers.BuscaTotalDePessoasCadastradas)
	r.Run(":5000") //listen and serve on localhost:5000
}
