package main

import (
	"markdown-version-control/application/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Routes(gin.Default()).Run()
}
