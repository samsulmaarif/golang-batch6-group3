package service

import (
	"database/sql"
	"golang-batch6-group3/adaptor"
	"golang-batch6-group3/server/params"
	"golang-batch6-group3/server/repository"
	"golang-batch6-group3/server/view"
	"log"
	"time"

	"github.com/google/uuid"
)

type TransactionServices struct {
	repo            repository.TransactionRepo
	typicodeAdaptor *adaptor.TypicodeAdaptor
}

func NewTransactionServices(repo repository.TransactionRepo, typicodeAdaptor *adaptor.TypicodeAdaptor) *TransactionServices {
	return &TransactionServices{
		repo:            repo,
		typicodeAdaptor: typicodeAdaptor,
	}
}

func (t *TransactionServices) GetTransactions() *view.Response {
	transactions, err := t.repo.GetTransactions()
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFindAll(view.NewTransactionFindAllResponse(transactions))
}

func (t *TransactionServices) GetMemberTransactions(id string) *view.Response {
	transactions, err := t.repo.GetMemberTransactions(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessFindAll(view.NewTransactionFindAllResponse(transactions))
}

func (t *TransactionServices) CreateTransaction(req *params.TransactionCreate) *view.Response {
	transaction := req.ParseToModel()

	transaction.Id = uuid.NewString()
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	err := t.repo.AddTransaction(transaction)
	if err != nil {
		log.Printf("get error add transaction with error %v\n", "")
		return view.ErrInternalServer(err.Error())
	}

	data, err := t.typicodeAdaptor.GetAllTypicode()
	if err != nil {
		return view.ErrInternalServer(err.Error())
	}

	return view.SuccessCreated(data)
}

func (u *TransactionServices) FindTransactionById(id string) *view.Response {
	transaction, err := u.repo.FindTransactionById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return view.ErrNotFound()
		}
		return view.ErrInternalServer(err.Error())
	}
	return view.SuccessFindAll(transaction)
}
