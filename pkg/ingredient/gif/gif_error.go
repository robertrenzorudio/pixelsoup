package gif

import (
	"fmt"
)

type ErrInputParameterValue struct {
	field  string
	reason string
}

func (m *ErrInputParameterValue) Error() string {
	return fmt.Sprintf("invalid VidToGifInput field: %s, reason: %s",
		m.field, m.reason)
}
