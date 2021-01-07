package period

import (
	"errors"
	"math"
	"time"

	"github.com/cyrildever/period/timestamp"
)

// The period module is a shared module, meaning once initialized you can't have different instance of it.
// In other words, its common parameters (begin of time, span) are set once and for all.
// Only test environment could be changed at will (through the `period.SetTestEnvironment()` function).

var (
	initialized        bool
	beginningTimestamp uint64
	span               uint64 = 10000
)

// Init must be called before anything, otherwise a panic may occur.
//
// It needs two mandatory arguments: an origin timestamp and a non-null period span (both in milliseconds),
// and an optional argument (passing it `true` if this is a test environment).
// It may throw an `AlreadyInitializedError` because the period module can't be initialized twice.
func Init(originTimestamp, periodSpan uint64, isTestEnvironment ...bool) error {
	if initialized {
		return NewAlreadyInitializedError()
	}
	if periodSpan == 0 {
		return errors.New("period span can't be null")
	}
	beginningTimestamp = originTimestamp
	span = periodSpan
	if len(isTestEnvironment) == 1 && isTestEnvironment[0] {
		timestamp.InTestEnvironment = true
	}
	initialized = true
	return nil
}

// ResetOriginTimestamp allows to force a new setting of the begin of the time.
// NB: Use with caution as all previous calculations might go wrong.
func ResetOriginTimestamp(newTimestamp uint64) {
	beginningTimestamp = newTimestamp
}

//--- TYPES

// Period ...
type Period struct {
	ID                   uint64
	startTimestampMillis uint64
	endTimestampMillis   uint64
}

//--- METHODS

// StartTimestampMillis ...
func (p *Period) StartTimestampMillis() uint64 {
	if !initialized || span == 0 {
		panic("period wasn't initialized")
	}
	if p.startTimestampMillis == 0 {
		elapsedFromBeginning := (p.ID - 1) * uint64(GetDuration().Seconds()*1000) // Genesis period is 1
		p.startTimestampMillis = beginningTimestamp + elapsedFromBeginning
	}
	return p.startTimestampMillis
}

// EndTimestampMillis ...
func (p *Period) EndTimestampMillis() uint64 {
	if !initialized || span == 0 {
		panic("period wasn't initialized")
	}
	if p.endTimestampMillis == 0 {
		elapsedFromBeginning := p.ID * uint64(GetDuration().Seconds()*1000)
		p.endTimestampMillis = beginningTimestamp + elapsedFromBeginning - 1 // Genesis period is 1
	}
	return p.endTimestampMillis
}

// Next ...
func (p *Period) Next() *Period {
	return &Period{ID: p.ID + 1}
}

//--- FUNCTIONS

// Current ...
func Current() (Period, error) {
	if !initialized || span == 0 {
		panic("period wasn't initialized")
	}
	currentTimestamp, err := timestamp.CurrentMillis()
	if err != nil {
		return Period{}, err
	}
	return Get(currentTimestamp)
}

//

// Get ...
func Get(timestampMillis uint64) (p Period, err error) {
	if !initialized || span == 0 {
		panic("period wasn't initialized")
	}
	if timestampMillis < beginningTimestamp {
		err = timestamp.NewInvalidTimestampError("timestamp is lower than begin of period")
		return
	}
	length := float64(GetDuration().Milliseconds())
	period := math.Floor(float64(timestampMillis-beginningTimestamp)/length) + 1 // First period is 1
	p = Period{ID: uint64(period)}
	return
}

// GetDuration ...
func GetDuration() time.Duration {
	if !initialized || span == 0 {
		panic("period wasn't initialized")
	}
	duration := span * 1e6
	return time.Duration(duration)
}

// Now returns the current period (with an ID of 0 if an error occurred)
func Now() Period {
	current, _ := Current()
	return current
}

// Span returns the period span (in milliseconds)
func Span() uint64 {
	return span
}

//--- utility

func IsTestEnvironment() bool {
	return timestamp.InTestEnvironment
}

func SetTestEnvironment(isTest bool) bool {
	timestamp.InTestEnvironment = isTest
	return IsTestEnvironment()
}

//-- ERRORS

// AlreadyInitializedError ...
type AlreadyInitializedError struct {
	message string
}

func (e *AlreadyInitializedError) Error() string {
	return e.message
}

// NewAlreadyInitializedError ...
func NewAlreadyInitializedError() *AlreadyInitializedError {
	return &AlreadyInitializedError{
		message: "period already initialized",
	}
}
