package router

import (
	c "golangMysql/controllers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) *gin.RouterGroup {
	v1 := r.Group("/v1")
	{
		v1.GET("/listar", c.ListProduto)
		v1.POST("/cadastro", c.CreateProduto)
		v1.PUT("/update/:id", c.UpdateProduto)
		v1.DELETE("/delete/:id", c.DeleteProduto)
	}
	return v1
}
