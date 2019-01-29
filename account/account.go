package account

import (
	"github.com/dogmatiq/dogma"
	"github.com/dogmatiq/example/messages"
)

// account is the aggregate root for a bank account.
type account struct {
	// Balance is the current account balance, in cents.
	Balance int64
}

func (a *account) ApplyEvent(m dogma.Message) {
	switch x := m.(type) {
	case messages.AccountCreditedForDeposit:
		a.Balance += x.Amount
	case messages.AccountCreditedForTransfer:
		a.Balance += x.Amount
	case messages.AccountDebitedForWithdrawal:
		a.Balance -= x.Amount
	case messages.AccountDebitedForTransfer:
		a.Balance -= x.Amount
	}
}

// Aggregate implements the business logic for a bank account.
//
// It centralizes all transactions that are applied to an account in order to
// enforce a strict no-overdraw policy.
type Aggregate struct{}

// New returns a new account instance.
func (Aggregate) New() dogma.AggregateRoot {
	return &account{}
}

// Configure configures the behavior of the engine as it relates to this
// handler.
func (Aggregate) Configure(c dogma.AggregateConfigurer) {
	c.Name("account")
	c.RouteCommandType(messages.OpenAccount{})
	c.RouteCommandType(messages.CreditAccountForDeposit{})
	c.RouteCommandType(messages.CreditAccountForTransfer{})
	c.RouteCommandType(messages.DebitAccountForWithdrawal{})
	c.RouteCommandType(messages.DebitAccountForTransfer{})
}

// RouteCommandToInstance returns the ID of the aggregate instance that is
// targetted by m.
func (Aggregate) RouteCommandToInstance(m dogma.Message) string {
	switch x := m.(type) {
	case messages.OpenAccount:
		return x.AccountID
	case messages.CreditAccountForDeposit:
		return x.AccountID
	case messages.CreditAccountForTransfer:
		return x.AccountID
	case messages.DebitAccountForWithdrawal:
		return x.AccountID
	case messages.DebitAccountForTransfer:
		return x.AccountID
	default:
		panic(dogma.UnexpectedMessage)
	}
}

// HandleCommand handles a command message that has been routed to this
// handler.
func (Aggregate) HandleCommand(s dogma.AggregateCommandScope, m dogma.Message) {
	switch x := m.(type) {
	case messages.OpenAccount:
		openAccount(s, x)
	case messages.CreditAccountForDeposit:
		creditForDeposit(s, x)
	case messages.CreditAccountForTransfer:
		creditForTransfer(s, x)
	case messages.DebitAccountForWithdrawal:
		debitForWithdrawal(s, x)
	case messages.DebitAccountForTransfer:
		debitForTransfer(s, x)
	default:
		panic(dogma.UnexpectedMessage)
	}
}

func openAccount(s dogma.AggregateCommandScope, m messages.OpenAccount) {
	if !s.Create() {
		s.Log("account has already been opened")
		return
	}

	s.RecordEvent(messages.AccountOpened{
		AccountID: m.AccountID,
		Name:      m.Name,
	})
}

func creditForDeposit(s dogma.AggregateCommandScope, m messages.CreditAccountForDeposit) {
	s.RecordEvent(messages.AccountCreditedForDeposit{
		TransactionID: m.TransactionID,
		AccountID:     m.AccountID,
		Amount:        m.Amount,
	})
}

func creditForTransfer(s dogma.AggregateCommandScope, m messages.CreditAccountForTransfer) {
	s.RecordEvent(messages.AccountCreditedForTransfer{
		TransactionID: m.TransactionID,
		AccountID:     m.AccountID,
		Amount:        m.Amount,
	})
}

func debitForWithdrawal(s dogma.AggregateCommandScope, m messages.DebitAccountForWithdrawal) {
	a := s.Root().(*account)

	if a.Balance >= m.Amount {
		s.RecordEvent(messages.AccountDebitedForWithdrawal{
			TransactionID: m.TransactionID,
			AccountID:     m.AccountID,
			Amount:        m.Amount,
		})
	} else {
		s.RecordEvent(messages.WithdrawalDeclined{
			TransactionID: m.TransactionID,
			AccountID:     m.AccountID,
			Amount:        m.Amount,
		})
	}
}

func debitForTransfer(s dogma.AggregateCommandScope, m messages.DebitAccountForTransfer) {
	a := s.Root().(*account)

	if a.Balance >= m.Amount {
		s.RecordEvent(messages.AccountDebitedForTransfer{
			TransactionID: m.TransactionID,
			AccountID:     m.AccountID,
			Amount:        m.Amount,
		})
	} else {
		s.RecordEvent(messages.TransferDeclined{
			TransactionID: m.TransactionID,
			AccountID:     m.AccountID,
			Amount:        m.Amount,
		})
	}
}