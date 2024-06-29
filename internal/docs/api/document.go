package api

import (
	"github.com/gin-gonic/gin"
	"github.com/karma-dev-team/karma-docs/internal/docs"
)

type DocumentHandler struct {
	documentService docs.DocumentService
}

func (h *DocumentHandler) GetDocument(c *gin.Context) {
	return 
}
