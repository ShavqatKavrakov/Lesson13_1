package wallet

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ShavqatKavrakov/Lesson13_1/pkg/types"
	"github.com/google/uuid"
)

type testService struct {
	*Service
}

func newTestService() *testService {
	return &testService{Service: &Service{}}
}

type testAccount struct {
	phone    types.Phone
	balance  types.Money
	payments []struct {
		amount   types.Money
		category types.PaymentCategory
	}
}

func (svc *testService) addAccount(data testAccount) (*types.Account, []*types.Payment, error) {
	account, err := svc.RegisterAccount(data.phone)
	if err != nil {
		return nil, nil, fmt.Errorf("can't register account, error= %v", err)
	}
	_, err = svc.Deposit(account.ID, data.balance)
	if err != nil {
		return nil, nil, fmt.Errorf("can't deposit account, error=%v", err)
	}
	payments := make([]*types.Payment, len(data.payments))
	for i, payment := range data.payments {
		payments[i], err = svc.Pay(account.ID, payment.category, payment.amount)
		if err != nil {
			return nil, nil, fmt.Errorf("can't make payment, error :=%v", err)
		}
	}
	return account, payments, nil
}

var defaultTestAccount = testAccount{
	phone:   "+992000001",
	balance: 10_000_00,
	payments: []struct {
		amount   types.Money
		category types.PaymentCategory
	}{
		{
			amount:   1_000_00,
			category: "auto",
		},
	},
}

func TestService_FindAccountById_success(t *testing.T) {
	svc := newTestService()
	account, err := svc.RegisterAccount("+99200000001")
	if err != nil {
		fmt.Println(err)
	}
	_, err = svc.FindAccountById(account.ID)
	if err == nil {
		fmt.Println(err)
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
	svc := newTestService()
	_, paymens, err := svc.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
	}
	payment := paymens[0]
	got, err := svc.FindPaymentById(payment.ID)
	if err != nil {
		t.Errorf("findPaymentById(): error=%v", err)
		return
	}
	if !reflect.DeepEqual(payment, got) {
		t.Errorf("FindPaymnetById():wrong payment returned =%v", err)
	}
}
func TestService_FindPaymentById_faid(t *testing.T) {
	svc := newTestService()
	_, _, err := svc.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = svc.FindPaymentById(uuid.New().String())
	if err == nil {
		t.Errorf("FindPaymentById():must return error, returned nil ")
		return
	}
	if err != ErrPaymentNotFound {
		t.Errorf("FindByPaymentId():must return ErrPaymentNotFound, returned =%v", err)
		return
	}
}
func TestService_Reject_success(t *testing.T) {
	svc := newTestService()
	_, payments, err := svc.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
	}
	payment := payments[0]
	err = svc.Reject(payment.ID)
	if err != nil {
		t.Error(err)
		return
	}
	savedPayment, err := svc.FindPaymentById(payment.ID)
	if err != nil {
		t.Errorf("Reject(): can't find payment by id, error=%v", err)
		return
	}
	if savedPayment.Status != types.PaymentStatusFail {
		t.Errorf("Reject(): status didn't changed,payment=%v", savedPayment)
		return
	}
	savedAccount, err := svc.FindAccountById(savedPayment.AccountID)
	if err != nil {
		t.Errorf("Reject():can't find account by id, error=%v", err)
		return
	}
	if savedAccount.Balance != defaultTestAccount.balance {
		t.Errorf("reject():balance didn't changed, account=%v", savedAccount)
	}
}
func TestService_Reject_faid(t *testing.T) {
	svc := newTestService()
	_, _, err := svc.addAccount(defaultTestAccount)
	if err != nil {
		t.Error(err)
		return
	}
	err = svc.Reject(uuid.New().String())
	if err == nil {
		t.Errorf("Reject():must return error, returned nil ")
		return
	}
}
