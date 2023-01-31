package handlers

import (
	"net/http"

	"github.com/aki-eiger/monorepo/module1"
	"github.com/gin-gonic/gin"
)

func Echo(c *gin.Context) {
	echostring := c.Query("echo")
	response := module1.CommonFunctionality(echostring)
	c.String(http.StatusOK, response)
}
