package repositories

import (
	"fmt"
	"housy/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	FindTransactionByUser(ID int) ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, orderId int) (models.Transaction, error)
	// UpdateFullcounter(transaction models.Transaction) (models.Transaction, error)
	// DeleteTransaction(transaction models.Transaction) (models.Transaction, error)
	// GetCategoryfilm(ID int) (models.Category, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("House").Preload("User").Find(&transactions).Error

	return transactions, err
}

func (r *repository) FindTransactionByUser(ID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("House").Preload("User").Find(&transactions, "user_id = ?", ID).Error
	fmt.Println(ID)
	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("House").Preload("User").First(&transaction, ID).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("House").Preload("User").Create(&transaction).Error

	return transaction, err
}
func (r *repository) UpdateTransaction(status string, orderId int) (models.Transaction, error) {
	var transaction models.Transaction
	r.db.Preload("User").Preload("House").First(&transaction, orderId)

	if status != transaction.Status && status == "success" {
		var house models.House
		r.db.First(&house, transaction.House.ID)
		house.Status = house.Status + transaction.Status
		// user.Subcribe = true
		r.db.Save(&house)
	}

	transaction.Status = status
	err := r.db.Save(&transaction).Error
	return transaction, err
}
