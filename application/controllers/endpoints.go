package controllers

import (
	"fmt"
	"markdown-version-control/application/model"
	"markdown-version-control/application/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var service = services.MarkdownService

// HealthCheck - HealthCheck
func HealthCheck(ctx *gin.Context) {
	payload := gin.H{
		"message":   "ping, pong!",
		"timestamp": time.Now().Unix(),
	}

	ctx.JSON(http.StatusOK, payload)
}

// FindAllMarkDowns - todo
func FindAllMarkDowns(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"data": service.FindAll()})
}

// FindMarkdownByID - todo
func FindMarkdownByID(ctx *gin.Context) {
	var id string = ctx.Param("projectId")

	if markdown, err := service.FindByID(id); err != nil {
		ctx.Status(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"data": markdown})
	}
}

// NewMarkdown - todo
func NewMarkdown(ctx *gin.Context) {
	var markdown model.Markdown

	if err := ctx.ShouldBindJSON(&markdown); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if contentID, err := service.Add(markdown); err == nil {
		ctx.Header("Location: ", fmt.Sprintf("%s/%s", ctx.FullPath(), contentID))
		ctx.Status(http.StatusCreated)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

// UpdateMarkdown - todo
func UpdateMarkdown(ctx *gin.Context) {
	var markdown model.Markdown

	if err := ctx.ShouldBindJSON(&markdown); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var id string = ctx.Param("projectId")

	if err := service.Update(id, markdown); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	} else {
		ctx.Status(http.StatusNoContent)
	}
}

// DeleteMarkdown - todo
func DeleteMarkdown(ctx *gin.Context) {
	var id string = ctx.Param("projectId")

	if err := service.Delete(id); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
	} else {
		ctx.Status(http.StatusNoContent)
	}
}
