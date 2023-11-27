// routes/routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"snakevalidator/game"
)

func DefineRoutes(r *gin.Engine) {
	r.GET("/new", game.NewGame)
	r.POST("/validate", game.ValidateMoves)
}
