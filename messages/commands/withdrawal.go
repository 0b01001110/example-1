package commands

import (
	"fmt"

	"github.com/dogmatiq/example/messages"
)

// Withdraw is a command requesting that funds be withdrawn from a bank account.
type Withdraw struct {
	TransactionID string
	AccountID     string
	Amount        int64
	ScheduledDate string
}

// ApproveWithdrawal is a command that approves an account withdrawal.
type ApproveWithdrawal struct {
	TransactionID string
	AccountID     string
	Amount        int64
}

// DeclineWithdrawal is a command that rejects an account withdrawal.
type DeclineWithdrawal struct {
	TransactionID string
	AccountID     string
	Amount        int64
	Reason        messages.DebitFailureReason
}

// String returns a human-readable description of the message.
func (m *Withdraw) String() string {
	return fmt.Sprintf(
		"withdraw %s from account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.AccountID),
	)
}

// String returns a human-readable description of the message.
func (m *ApproveWithdrawal) String() string {
	return fmt.Sprintf(
		"approve withdrawal of %s from account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.AccountID),
	)
}

// String returns a human-readable description of the message.
func (m *DeclineWithdrawal) String() string {
	return fmt.Sprintf(
		"decline withdrawal of %s from account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.AccountID),
	)
}
