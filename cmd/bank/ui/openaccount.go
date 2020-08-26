package ui

import (
	"fmt"
	"math/rand"
)

// generateAccountID returns a new unique customer ID.
func generateAccountID() string {
	return fmt.Sprintf(
		"%02d-%03d-%03d",
		10+rand.Intn(99),
		100+rand.Intn(899),
		100+rand.Intn(899),
	)
}
