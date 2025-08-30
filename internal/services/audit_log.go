package services

import (
	"time"

	"github.com/Zekeriyyah/stellar-x/internal/models"
	"github.com/Zekeriyyah/stellar-x/internal/repositories"
)

type AuditLogService struct {
	AuditRepo *repositories.AuditLogRepository
}

func NewAuditLogService(auditRepo *repositories.AuditLogRepository) *AuditLogService {
	return &AuditLogService{
		AuditRepo: auditRepo,
	}
}

// LogRequest creates an audit log entry
func (s *AuditLogService) LogRequest(userID uint, walletID uint, ip, device, browser, country, endpoint, method string) error {
	log := &models.AuditLog{
		UserID:    userID,
		WalletID:  &walletID,
		IPAddress: ip,
		Device:    device,
		Browser:   browser,
		Country:   country,
		Path:	   endpoint,
		Method:	   method,
		CreatedAt: time.Now(),
	}

	return s.AuditRepo.Create(log)
}

func (s *AuditLogService) GetAllLogs(userId uint) ([]models.AuditLog, error) {
	logs, err := s.AuditRepo.GetAuditLogByUserID(userId)
	if err != nil {
		return []models.AuditLog{}, err
	}
	return logs, nil
}