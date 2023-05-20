package types

import (
	"testing"
	"time"
)

func TestTimestamp_AddSeconds(t *testing.T) {
	timestamp := NewTimestamp()

	seconds := 5

	expected := time.Time(timestamp).Add(time.Second * time.Duration(seconds))

	newTimestamp := timestamp.AddSeconds(seconds)

	if !time.Time(newTimestamp).Equal(expected) {
		t.Errorf("Expected %v, got %v", expected, time.Time(newTimestamp))
	}
}

func TestTimestamp_IsBefore(t *testing.T) {
	t1 := NewTimestamp()

	time.Sleep(time.Millisecond * 10)

	t2 := NewTimestamp()

	isBefore := t1.IsBefore(t2)

	if !isBefore {
		t.Errorf("Expected %v, got %v", true, isBefore)
	}
}
