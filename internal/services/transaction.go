package services

import (
	"math"

	"github.com/Zekeriyyah/stellar-x/internal/models"
	"github.com/Zekeriyyah/stellar-x/internal/repositories"
)

type TransactionService struct {
	TransactionRepo *repositories.TransactionRepository
	WalletRepo      *repositories.WalletRepository
}

func NewTransactionService(txRepo *repositories.TransactionRepository, walletRepo *repositories.WalletRepository) *TransactionService {
	return &TransactionService{
		TransactionRepo: txRepo,
		WalletRepo:      walletRepo,
	}
}

// GetByUserID retrieves all transactions for a user
func (s *TransactionService) GetByUserID(userID uint) ([]models.Transaction, error) {
	return s.TransactionRepo.GetTransactionsByUserID(userID)
}

// FOR TRANSACTION EXPLORER
// GetRecent retrieves paginated recent transactions with metadata
func (s *TransactionService) GetRecent(page, pageSize int) (*PaginatedTransactions, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100 // Limit max page size
	}

	offset := (page - 1) * pageSize

	transactions, err := s.TransactionRepo.GetRecent(pageSize, offset)
	if err != nil {
		return nil, err
	}

	total, err := s.TransactionRepo.GetCount()
	if err != nil {
		return nil, err
	}

	return &PaginatedTransactions{
		Data:       transactions,
		Page:       page,
		PageSize:   pageSize,
		Total:      total,
		TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
	}, nil
}

// PaginatedTransactions holds paginated response
type PaginatedTransactions struct {
	Data       []models.Transaction `json:"data"`
	Page       int                  `json:"page"`
	PageSize   int                  `json:"page_size"`
	Total      int64                `json:"total"`
	TotalPages int                  `json:"total_pages"`
}