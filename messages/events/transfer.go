package events

import (
	"fmt"

	"github.com/dogmatiq/example/messages"
)

// TransferStarted is an event indicating that the process of transferring funds
// from one account to another has begun.
type TransferStarted struct {
	TransactionID string
	FromAccountID string
	ToAccountID   string
	Amount        int64
	ScheduledDate string
}

// TransferApproved is an event that indicates a requested transfer has been
// approved.
type TransferApproved struct {
	TransactionID string
	FromAccountID string
	ToAccountID   string
	Amount        int64
}

// TransferDeclined is an event that indicates a requested transfer has been
// declined.
type TransferDeclined struct {
	TransactionID string
	FromAccountID string
	ToAccountID   string
	Amount        int64
	Reason        messages.DebitFailureReason
}

// String returns a human-readable description of the message.
func (m *TransferStarted) String() string {
	return fmt.Sprintf(
		"transfer %s from account %s to account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.FromAccountID),
		messages.FormatID(m.ToAccountID),
	)
}

// String returns a human-readable description of the message.
func (m *TransferApproved) String() string {
	return fmt.Sprintf(
		"transfer approved for %s from account %s to account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.FromAccountID),
		messages.FormatID(m.ToAccountID),
	)
}

// String returns a human-readable description of the message.
func (m *TransferDeclined) String() string {
	return fmt.Sprintf(
		"transfer declined for %s from account %s to account %s",
		messages.FormatAmount(m.Amount),
		messages.FormatID(m.FromAccountID),
		messages.FormatID(m.ToAccountID),
	)
}
