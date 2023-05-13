package types

import "time"

type Timestamp time.Time

func NewTimestamp() Timestamp {
	return Timestamp(time.Now().UTC())
}
