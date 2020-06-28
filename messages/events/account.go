package events

import (
	"fmt"

	"github.com/dogmatiq/example/messages"
)

// AccountOpened is an event indicating that a new bank account has been opened.
type AccountOpened struct {
	CustomerID  string
	AccountID   string
	AccountName string
}

// AccountCredited is an event indicating that a bank account was credited.
type AccountCredited struct {
	TransactionID   string
	AccountID       string
	TransactionType messages.TransactionType
	Amount          int64
}

// AccountDebited is an event indicating that a bank account was debited.
type AccountDebited struct {
	TransactionID   string
	AccountID       string
	TransactionType messages.TransactionType
	Amount          int64
	ScheduledDate   string
}

// AccountDebitDeclined is an event indicating that a bank account debit was
// declined.
type AccountDebitDeclined struct {
	TransactionID   string
	AccountID       string
	TransactionType messages.TransactionType
	Amount          int64
	Reason          messages.DebitFailureReason
}

// String returns a human-readable description of the message.
func (m *AccountOpened) String() string {
	return fmt.Sprintf(
		"account %s %s opened for customer %s",
		messages.FormatID(m.AccountID),
		m.AccountName,
		messages.FormatID(m.CustomerID),
	)
}

// String returns a human-readable description of the message.
func (m *AccountCredited) String() string {
	return fmt.Sprintf(
		"credited %s to account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.AccountID),
	)
}

// String returns a human-readable description of the message.
func (m *AccountDebited) String() string {
	return fmt.Sprintf(
		"debited %s from account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.AccountID),
	)
}

// String returns a human-readable description of the message.
func (m *AccountDebitDeclined) String() string {
	return fmt.Sprintf(
		"declined debit of %s from account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.AccountID),
	)
}
