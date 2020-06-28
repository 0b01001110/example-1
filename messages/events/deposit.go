package events

import (
	"fmt"

	"github.com/dogmatiq/example/messages"
)

// DepositStarted is an event indicating that the process of depositing funds
// into an account has begun.
type DepositStarted struct {
	TransactionID string
	AccountID     string
	Amount        int64
}

// DepositApproved is an event that indicates a requested deposit has been
// approved.
type DepositApproved struct {
	TransactionID string
	AccountID     string
	Amount        int64
}

// String returns a human-readable description of the message.
func (m *DepositStarted) String() string {
	return fmt.Sprintf(
		"depositing %s into account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.AccountID),
	)
}

// String returns a human-readable description of the message.
func (m *DepositApproved) String() string {
	return fmt.Sprintf(
		"deposit approved of %s into account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.AccountID),
	)
}
