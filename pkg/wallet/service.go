package wallet

import (
	"errors"

	"github.com/ShavqatKavrakov/Lesson13_1/pkg/types"
	"github.com/google/uuid"
)

var ErrPhoneRegistered = errors.New("phone already registered")
var ErrAccountNotFound = errors.New("account not found")
var ErrAmountMostBePositive = errors.New("amount must be greater than zero")
var ErrNotEnouthBalance = errors.New("not enough balance in account")
var ErrPaymentNotFound = errors.New("payment not found")

type Service struct {
	nextAccountId int64
	accounts      []*types.Account
	payments      []*types.Payment
}

func (s *Service) RegisterAccount(phone types.Phone) (*types.Account, error) {
	for _, acc := range s.accounts {
		if acc.Phone == phone {
			return nil, ErrPhoneRegistered
		}
	}
	s.nextAccountId++
	account := &types.Account{
		ID:      s.nextAccountId,
		Phone:   phone,
		Balance: 0,
	}
	s.accounts = append(s.accounts, account)
	return account, nil
}

func (s *Service) FindAccountById(accountId int64) (*types.Account, error) {
	for _, account := range s.accounts {
		if account.ID == accountId {
			return account, nil
		}
	}
	return nil, ErrAccountNotFound
}

func (s *Service) Deposit(accountId int64, amount types.Money) (*types.Account, error) {
	if amount <= 0 {
		return nil, ErrAmountMostBePositive
	}
	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == accountId {
			account = acc
			break
		}
	}
	if account == nil {
		return nil, ErrAccountNotFound
	}
	account.Balance += amount
	return account, nil
}

func (s *Service) Pay(acountId int64, category types.PaymentCategory, amount types.Money) (*types.Payment, error) {
	if amount < 0 {
		return nil, ErrAmountMostBePositive
	}
	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == acountId {
			account = acc
			break
		}
	}
	if account == nil {
		return nil, ErrAccountNotFound
	}
	if account.Balance < amount {
		return nil, ErrNotEnouthBalance
	}
	account.Balance -= amount

	paymentID := uuid.New().String()
	payment := &types.Payment{
		ID:        paymentID,
		AccountID: acountId,
		Amount:    amount,
		Category:  category,
		Status:    types.PaymentStatusInProgress,
	}
	s.payments = append(s.payments, payment)
	return payment, nil
}

func (s *Service) FindPaymentById(paymentId string) (*types.Payment, error) {
	for _, payment := range s.payments {
		if payment.ID == paymentId {
			return payment, nil
		}
	}
	return nil, ErrPaymentNotFound
}
func (s *Service) Reject(paymentId string) error {
	payment, err := s.FindPaymentById(paymentId)
	if err != nil {
		return err
	}
	payment.Status = types.PaymentStatusFail
	var account *types.Account
	for _, acc := range s.accounts {
		if acc.ID == payment.AccountID {
			account = acc
			break
		}
	}
	account.Balance += payment.Amount
	return nil
}
