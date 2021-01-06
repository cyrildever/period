package timestamp

import (
	"time"

	"github.com/cyrildever/go-utls/common/ntp"
)

// Set to `true` if NTP can't be initialized before testing, other you'll get an InvalidTimestampError.
var InTestEnvironment bool

// CurrentMillis ...
func CurrentMillis() (uint64, error) {
	if InTestEnvironment {
		return uint64(time.Now().UnixNano() / 1e6), nil
	}
	time, ntpErr := ntp.Time("")
	if ntpErr != nil {
		return 0, NewInvalidTimestampError("ntp.Time has failed: timestamp can't be returned")
	}
	milliseconds := time.UnixNano() / 1e6
	return uint64(milliseconds), nil
}

// InvalidTimestampError ...
type InvalidTimestampError struct {
	message string
}

func (e *InvalidTimestampError) Error() string {
	return e.message
}

// NewInvalidTimestampError ...
func NewInvalidTimestampError(message string) *InvalidTimestampError {
	return &InvalidTimestampError{
		message: message,
	}
}
