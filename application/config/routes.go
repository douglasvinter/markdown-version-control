package config

import (
	"markdown-version-control/application/controllers"

	"github.com/gin-gonic/gin"
)

// Routes - HTTP routes contaning the endpoints paths & verbs
var Routes = func(app *gin.Engine) *gin.Engine {
	app.GET("/health-check", controllers.HealthCheck)

	base := app.Group("/api/markdown")

	base.GET("/web-hook", controllers.FindAllMarkDowns)

	base.GET("/web-hook/:projectId", controllers.FindMarkdownByID)

	base.POST("/web-hook", controllers.NewMarkdown)

	base.PUT("/web-hook/:projectId", controllers.UpdateMarkdown)

	base.DELETE("/web-hook/:projectId", controllers.DeleteMarkdown)

	return app
}
