package main

import (
	"golangMysql/router"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()
	router.Router(app)
	app.Run("localhost:8080")

}
