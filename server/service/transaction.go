package service

import (
	"database/sql"
	"golang-batch6-group3/server/params"
	"golang-batch6-group3/server/repository"
	"golang-batch6-group3/server/view"
	"log"
	"time"

	"github.com/google/uuid"
)

type TransactionServices struct {
	repo repository.TransactionRepo
}

func NewTransactionServices(repo repository.TransactionRepo) *TransactionServices {
	return &TransactionServices{
		repo: repo,
	}
}

func (t *TransactionServices) GetTransactions() *view.Response {
	transactions, err := t.repo.GetTransactions()
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("GET_TRANSACTION_FAIL", err.Error())
	}

	return view.SuccessFindAll(view.NewTransactionFindAllResponse(transactions))
}

func (t *TransactionServices) GetMemberTransactions(id string) *view.Response {
	transactions, err := t.repo.GetMemberTransactions(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("GET_MEMBER_TRANSACTION_FAIL", err.Error())
	}

	return view.SuccessFindAll(view.NewTransactionFindAllResponse(transactions))
}

func (t *TransactionServices) CreateTransaction(req *params.TransactionCreate) *view.Response {
	transaction := req.ParseToModel()

	transaction.Id = uuid.NewString()
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()
	transaction.Status = "pending"

	err := t.repo.AddTransaction(transaction)
	if err != nil {
		log.Printf("get error add transaction with error %v\n", "")
		return view.ErrInternalServer("CREATE_TRANSACTION_FAIL", err.Error())
	}

	return view.SuccessCreated(transaction)
}

func (t *TransactionServices) UpdateTransactionStatusById(id string, status string) *view.Response {
	err := t.repo.UpdateTransactionStatusById(id, status)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("UPDATE_TRANSACTION_FAIL", err.Error())
	}
	return view.SuccessFindAll("UPDATE_PRODUCT_SUCCESS")
}

func (u *TransactionServices) FindTransactionById(id string) *view.Response {
	transaction, err := u.repo.FindTransactionById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("FIND_TRANSACTION_FAIL", err.Error())
	}
	return view.SuccessFindAll(transaction)
}
