package events

import (
	"fmt"

	"github.com/dogmatiq/example/messages"
)

// WithdrawalStarted is an event indicating that the process of withdrawing
// funds from an account has begun.
type WithdrawalStarted struct {
	TransactionID string
	AccountID     string
	Amount        int64
	ScheduledDate string
}

// WithdrawalApproved is an event that indicates a requested withdrawal has been
// approved.
type WithdrawalApproved struct {
	TransactionID string
	AccountID     string
	Amount        int64
}

// WithdrawalDeclined is an event that indicates a requested withdrawal has been
// declined.
type WithdrawalDeclined struct {
	TransactionID string
	AccountID     string
	Amount        int64
	Reason        messages.DebitFailureReason
}

// String returns a human-readable description of the message.
func (m *WithdrawalStarted) String() string {
	return fmt.Sprintf(
		"withdrawing %s from account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.AccountID),
	)
}

// String returns a human-readable description of the message.
func (m *WithdrawalApproved) String() string {
	return fmt.Sprintf(
		"withdrawal approved for %s from account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.AccountID),
	)
}
