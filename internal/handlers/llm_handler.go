package handlers

import (
	"net/http"

	"github.com/Zekeriyyah/stellar-x/internal/services"
	"github.com/gin-gonic/gin"
)

type AIHandler struct {
	AIService *services.AIService
}

func NewAIHandler(aiService *services.AIService) *AIHandler {
	return &AIHandler{AIService: aiService}
}

// Ask handles GET /api/v1/ask?q=...
func (h *AIHandler) Ask(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Query parameter 'q' is required"})
		return
	}

	if h.AIService == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "AI service not configured"})
		return
	}

	answer, err := h.AIService.Ask(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"query":   query,
		"answer":  answer,
	})
}