package api

import (
	"github.com/gin-gonic/gin"
	"github.com/karma-dev-team/karma-docs/internal/docs"
)

func RegisterDocs(api *gin.RouterGroup, docsService docs.DocumentService) {
	docsHandler := &DocumentHandler{documentService: docsService}

	api.GET("/docs/:documentid", docsHandler.GetDocument)
}
