package entities

import "time"

type User struct {
	ID        uint64
	FirstName string
	LastName  string
	Email     string
	Age       uint8
	City      string
	Gender    Gender
	CreatedAt time.Time
	Surveys   []uint64
}

type Gender string

const (
	Unknown Gender = "unknown"
	Male    Gender = "male"
	Female  Gender = "female"
)
