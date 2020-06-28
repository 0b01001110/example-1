package commands

import (
	"fmt"

	"github.com/dogmatiq/example/messages"
)

// OpenAccountForNewCustomer is a command requesting that a new bank account be
// opened for a new customer.
type OpenAccountForNewCustomer struct {
	CustomerID   string
	CustomerName string
	AccountID    string
	AccountName  string
}

// OpenAccount is a command requesting that a new bank account be opened for an
// existing customer.
type OpenAccount struct {
	CustomerID  string
	AccountID   string
	AccountName string
}

// CreditAccount is a command that requests a bank account be credited.
type CreditAccount struct {
	TransactionID   string
	AccountID       string
	TransactionType messages.TransactionType
	Amount          int64
}

// DebitAccount is a command that requests a bank account be debited.
type DebitAccount struct {
	TransactionID   string
	AccountID       string
	TransactionType messages.TransactionType
	Amount          int64
	ScheduledDate   string
}

// String returns a human-readable description of the message.
func (m *OpenAccountForNewCustomer) String() string {
	return fmt.Sprintf(
		"account %s %s opened for new customer %s %s",
		messages.FormatID(m.AccountID),
		m.AccountName,
		messages.FormatID(m.CustomerID),
		m.CustomerName,
	)
}

// String returns a human-readable description of the message.
func (m *OpenAccount) String() string {
	return fmt.Sprintf(
		"account %s %s opened for customer %s",
		messages.FormatID(m.AccountID),
		m.AccountName,
		messages.FormatID(m.CustomerID),
	)
}

// String returns a human-readable description of the message.
func (m *CreditAccount) String() string {
	return fmt.Sprintf(
		"credit %s to account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.AccountID),
	)
}

// String returns a human-readable description of the message.
func (m *DebitAccount) String() string {
	return fmt.Sprintf(
		"debit %s from account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.AccountID),
	)
}
