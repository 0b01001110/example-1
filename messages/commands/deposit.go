package commands

import (
	"fmt"

	"github.com/dogmatiq/example/messages"
)

// Deposit is a command requesting that funds be deposited into a bank account.
type Deposit struct {
	TransactionID string
	AccountID     string
	Amount        int64
}

// ApproveDeposit is a command that approves an account deposit.
type ApproveDeposit struct {
	TransactionID string
	AccountID     string
	Amount        int64
}

// String returns a human-readable description of the message.
func (m *Deposit) String() string {
	return fmt.Sprintf(
		"deposit %s into account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.AccountID),
	)
}

// String returns a human-readable description of the message.
func (m *ApproveDeposit) String() string {
	return fmt.Sprintf(
		"approve deposit of %s into account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.AccountID),
	)
}
