package wallet

import (
	"fmt"
	"testing"
)

func TestService_FindAccountById_success(t *testing.T) {
	svc := &Service{}
	account, err := svc.RegisterAccount("+99200000001")
	if err != nil {
		fmt.Println(err)
	}
	_, err = svc.FindAccountById(account.ID)
	if err == nil {
		fmt.Println()
	}
}
func TestService_FindAccountById_notFound(t *testing.T) {
	svc := &Service{}
	account, err := svc.RegisterAccount("+99200000001")
	if err != nil {
		fmt.Println(err)
	}
	_, err = svc.FindAccountById(account.ID + 1)
	if err != nil {
		fmt.Println(err)
	}
}
func TestService_FindPaymentById_success(t *testing.T) {
	svc := &Service{}
	account, err := svc.RegisterAccount("+99200000001")
	if err != nil {
		fmt.Println(err)
	}
	svc.Deposit(account.ID, 100)
	payment, err := svc.Pay(account.ID, "auto", 50)
	if err != nil {
		fmt.Println(err)
	}
	_, err = svc.FindPaymentById(payment.ID)
	if err != nil {
		fmt.Println(err)
	}
}
func TestService_FindPaymentById_notFound(t *testing.T) {
	svc := &Service{}
	account, err := svc.RegisterAccount("+99200000001")
	if err != nil {
		fmt.Println(err)
	}
	svc.Deposit(account.ID, 100)
	payment, err := svc.Pay(account.ID, "auto", 50)
	if err != nil {
		fmt.Println(err)
	}
	_, err = svc.FindPaymentById(payment.ID + "1")
	if err != nil {
		fmt.Println(err)
	}
}
func TestService_Reject_success(t *testing.T) {
	svc := &Service{}
	account, err := svc.RegisterAccount("+99200000001")
	if err != nil {
		fmt.Println(err)
	}
	svc.Deposit(account.ID, 100)
	payment, err := svc.Pay(account.ID, "auto", 50)
	if err != nil {
		fmt.Println(err)
	}
	err = svc.Reject(payment.ID)
	if err != nil {
		fmt.Println()
	}
}
func TestService_Reject_notFound(t *testing.T) {
	svc := &Service{}
	account, err := svc.RegisterAccount("+99200000001")
	if err != nil {
		fmt.Println(err)
	}
	svc.Deposit(account.ID, 100)
	payment, err := svc.Pay(account.ID, "auto", 50)
	if err != nil {
		fmt.Println(err)
	}
	err = svc.Reject(payment.ID + "1")
	if err != nil {
		fmt.Println(err)
	}
}
