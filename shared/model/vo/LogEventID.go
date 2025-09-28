package vo

import (
	"fmt"
)

type LogEventID string

func NewLogEventID(random10Char string) LogEventID {
	return LogEventID(fmt.Sprintf("LOG-%s", random10Char))
}
