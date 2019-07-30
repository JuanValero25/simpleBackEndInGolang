package repository

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
	"time"
)

type TransactionType string

const (
	Debit  TransactionType = "debit"
	Credit TransactionType = "credit"
)

type Transaction struct {
	ID            string          `json:"id"`
	Type          TransactionType `json:"type"`
	Amount        float64         `json:"amount"`
	EffectiveDate time.Time       `json:"effectiveDate"`
}

type TrasactionError struct {
	Message string `json:"Message"`
}

func (t *TrasactionError) Error() string {
	return "Error Message : " + t.Message
}

var (
	mutex             sync.RWMutex
	mockedDBMap       = make(map[string]*Transaction)
	allTransaction    []*Transaction
	sumAllTransaction float64
)

type MockRepository struct {
}

func (c *MockRepository) GetTransactionById(trasanctionId string) (*Transaction, error) {
	if !IsValidUUID(trasanctionId) {
		return nil, &TrasactionError{"id is not valid"}
	}
	mutex.RLock()
	defer mutex.RUnlock()
	transactionFound := mockedDBMap[trasanctionId]
	return transactionFound, nil
}

func (c *MockRepository) GetAllTransaction() *[]*Transaction {
	mutex.RLock()
	defer mutex.RUnlock()
	return &allTransaction
}

func (c *MockRepository) GetBalance() *float64 {
	mutex.RLock()
	defer mutex.RUnlock()
	return &sumAllTransaction
}

func (c *MockRepository) PostTransaction(transaction *Transaction) error {
	err := IsInValidTransaction(transaction)
	if err != nil {
		return err
	}
	mutex.Lock()
	if transaction.Type == Debit && sumAllTransaction < transaction.Amount {
		return &TrasactionError{Message: fmt.Sprintf("Debit is to hight for current balance : %f", sumAllTransaction)}
	}
	defer mutex.Unlock()
	transaction.EffectiveDate = time.Now()
	transaction.ID = uuid.New().String()
	mockedDBMap[transaction.ID] = transaction
	allTransaction = append(allTransaction, transaction)
	SumTransaction(transaction)
	return nil
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func SumTransaction(transaction *Transaction) {
	switch transaction.Type {
	case Debit:
		sumAllTransaction = sumAllTransaction - transaction.Amount
	case Credit:
		sumAllTransaction = sumAllTransaction + transaction.Amount
	}
}

func IsInValidTransaction(transactionInsert *Transaction) *TrasactionError {
	if transactionInsert.Amount <= 0 {
		return &TrasactionError{Message: "amount can be 0 or negative"}
	}

	if transactionInsert.Type != Credit && transactionInsert.Type != Debit {
		return &TrasactionError{Message: "transactionType cant be diferente from debit and credit"}
	}
	return nil
}
