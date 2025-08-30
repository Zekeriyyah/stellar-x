package middleware

import (
	"github.com/Zekeriyyah/stellar-x/internal/services"
	"github.com/Zekeriyyah/stellar-x/pkg"
	"github.com/gin-gonic/gin"
)

const AuditInfoKey = "audit_info"
type AuditInfo struct {
	UserID		uint
	WalletID	uint
	IPAddress 	string
    Browser 	string
    Device  	string
    OS        	string
    Country   	string
    Path     	string
    Method    	string
}

func AuditLogger(auditService *services.AuditLogService) gin.HandlerFunc {
    return func(c *gin.Context) {
       
        // IP Address
        ip := c.ClientIP()

        ua := c.GetHeader("User-Agent")

        // Parse browser, OS, device
        userInfo := pkg.ParseUserAgent(ua)

        // get country from ipapi.co
        country := pkg.GetCountryFromIP(ip) 

        // Attach to context
        auditInfo := AuditInfo{
			IPAddress: ip,
            Browser:   userInfo.Browser,
            Device:    userInfo.Device,
            OS:        userInfo.OS,
            Country:   country,
            Path:      c.Request.URL.Path,
            Method:    c.Request.Method,
        }

        userId, exists := c.Get("user_id")
        if exists {
            auditInfo.UserID = userId.(uint)
        }

        c.Set(AuditInfoKey, auditInfo)
        c.Next()

		// After handler runs, log the request
		if auditService != nil {
			_ = auditService.LogRequest(
				auditInfo.UserID,
				auditInfo.WalletID,
				auditInfo.IPAddress,
				auditInfo.Device,
				auditInfo.Browser,
				auditInfo.Country,
				auditInfo.Path,
				auditInfo.Method)
		}
    }
}