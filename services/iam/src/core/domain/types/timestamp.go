package types

import "time"

type Timestamp time.Time

func NewTimestamp() Timestamp {
	return Timestamp(time.Now().UTC())
}

func (t Timestamp) AddSeconds(seconds int) Timestamp {
	return Timestamp(time.Time(t).Add(time.Second * time.Duration(seconds)))
}

func (t Timestamp) IsBefore(other Timestamp) bool {
	return time.Time(t).Before(time.Time(other))
}
