package domain

import "time"

type AutoIncr struct {
	ID        uint64
	CreatedAt time.Time `db:"created_at"`
}
