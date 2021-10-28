package entity

import "time"

type Customer struct {
	ID        int
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
