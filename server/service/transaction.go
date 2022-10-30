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
	repo     repository.TransactionRepo
	repoProd repository.ProductRepo
}

func NewTransactionServices(repo repository.TransactionRepo, repoProd repository.ProductRepo) *TransactionServices {
	return &TransactionServices{
		repo:     repo,
		repoProd: repoProd,
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

	return view.SuccessFindAll("GET_TRANSACTION_SUCCESS", view.NewTransactionFindAllResponse(transactions))
}

func (t *TransactionServices) GetMemberTransactions(id string) *view.Response {
	transactions, err := t.repo.GetMemberTransactions(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("GET_TRANSACTION_FAIL", err.Error())
	}

	return view.SuccessFindAll("GET_TRANSACTION_SUCCESS", view.NewTransactionFindAllResponse(transactions))
}

func (t *TransactionServices) CreateTransaction(req *params.TransactionCreate) *view.Response {
	transaction := req.ParseToModel()
	product, err := t.repoProd.FindProductById(transaction.ProductId)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("GET_PRODUCT_FAIL", err.Error())
	}

	transaction.Id = uuid.NewString()
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()
	transaction.Total = (req.Quantity * req.Ongkir) + (req.Quantity * product.Price)
	transaction.Status = "pending"

	err = t.repo.AddTransaction(transaction)
	if err != nil {
		log.Printf("get error add transaction with error %v\n", "")
		return view.ErrInternalServer("CREATE_TRANSACTION_FAIL", err.Error())
	}

	return view.SuccessAdd("ADD_TRANSACTION_SUCCESS", transaction)
}

func (t *TransactionServices) UpdateTransactionStatusById(id string, status string) *view.Response {
	err := t.repo.UpdateTransactionStatusById(id, status)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("UPDATE_TRANSACTION_FAIL", err.Error())
	}
	return view.SuccessUpdated("UPDATE_TRANSACTION_SUCCESS", "UPDATE_TRANSACTION_SUCCESS")
}

func (u *TransactionServices) FindTransactionById(id string) *view.Response {
	transaction, err := u.repo.FindTransactionById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer("FIND_TRANSACTION_FAIL", err.Error())
	}
	return view.SuccessFindAll("FIND_TRANSACTION_SUCCESS", transaction)
}
