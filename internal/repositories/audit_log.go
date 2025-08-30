package repositories

import (
	"github.com/Zekeriyyah/stellar-x/internal/models"
	"gorm.io/gorm"
)


type AuditLogRepository struct {
	DB *gorm.DB
}

func NewAuditLogRepository(db *gorm.DB) *AuditLogRepository {
	return &AuditLogRepository{
		DB: db,
	}
}

func (r *AuditLogRepository) Create(log *models.AuditLog) error {
	return r.DB.Create(log).Error
}

func (r *AuditLogRepository) GetAuditLogByUserID(userID uint) ([]models.AuditLog, error) {
	var logs []models.AuditLog
	err := r.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&logs).Error
	return logs, err
}