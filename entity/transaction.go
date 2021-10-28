package entity

import "time"

type Transaction struct {
	ID        int
	Code      string
	Amount    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
