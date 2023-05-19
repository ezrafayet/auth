package types

import (
	"testing"
	"time"
)

func TestTimestamp_NewTimestamp(t *testing.T) {
	now := time.Now().UTC()
	timestamp := NewTimestamp()

	if time.Time(timestamp).Before(now) {
		t.Errorf("NewTimestamp() = %v is before now = %v", timestamp, now)
	}
}

func TestTimestamp_AddSeconds(t *testing.T) {
	timestamp := NewTimestamp()
	seconds := 5
	expected := time.Time(timestamp).Add(time.Second * time.Duration(seconds))

	newTimestamp := timestamp.AddSeconds(seconds)
	if !time.Time(newTimestamp).Equal(expected) {
		t.Errorf("AddSeconds() = %v, want %v", newTimestamp, expected)
	}
}

func TestTimestamp_IsBefore(t *testing.T) {
	t1 := NewTimestamp()
	time.Sleep(time.Second)
	t2 := NewTimestamp()

	if !t1.IsBefore(t2) {
		t.Errorf("Expected t1 to be before t2, but it was not")
	}
}
