package handlers

import (
	"net/http"
	"strconv"

	"github.com/Zekeriyyah/stellar-x/internal/services"
	"github.com/gin-gonic/gin"
)

type AuditLogHandler struct {
	AuditService *services.AuditLogService
}

func NewAuditLogHandler(auditService *services.AuditLogService) *AuditLogHandler {
	return &AuditLogHandler{AuditService: auditService}
}

// GetLogs handles GET /api/v1/audit/:userId
func (h *AuditLogHandler) GetAuditLogByUserID(c *gin.Context) {
	userID := c.Param("userId")

	// parse user id
		userId, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID: must be a positive number"})
		return
	}
	userIdUint := uint(userId)

	logs, err := h.AuditService.GetAllLogs(userIdUint)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no audit logs found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userId": userID,
		"logs":   logs,
	})
}