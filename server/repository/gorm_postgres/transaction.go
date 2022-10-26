package gorm_postgres

import (
	"golang-batch6-group3/server/model"
	"golang-batch6-group3/server/repository"

	"gorm.io/gorm"
)

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepoGormPostgres(db *gorm.DB) repository.TransactionRepo {
	return &transactionRepo{
		db: db,
	}
}

func (t *transactionRepo) GetTransactions() (*[]model.Transaction, error) {
	var transactions []model.Transaction
	err := t.db.Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return &transactions, nil
}

func (t *transactionRepo) GetMemberTransactions(id string) (*[]model.Transaction, error) {
	var transactions []model.Transaction
	err := t.db.Where("user_id=?", id).Find(&transactions).Error
	if err != nil {
		return nil, err
	}
	return &transactions, nil
}

func (t *transactionRepo) AddTransaction(transaction *model.Transaction) error {
	err := t.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (t *transactionRepo) UpdateTransactionStatusById(id string, status string) error {
	err := t.db.Model(model.Transaction{}).Where("id=?", id).Updates(model.Transaction{Status: status}).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *transactionRepo) FindTransactionById(id string) (*model.Transaction, error) {
	var transaction model.Transaction
	err := t.db.Where("id=?", id).First(&transaction).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
