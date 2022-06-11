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
	_, err := svc.RegisterAccount("+99200000001")
	if err != nil {
		fmt.Println(err)
	}
	_, err = svc.FindAccountById(2)
	if err != nil {
		fmt.Println(err)
	}
}
