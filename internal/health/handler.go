package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(c *gin.Context) {
	response := GetHealth()
	c.JSON(http.StatusOK, response)
}
