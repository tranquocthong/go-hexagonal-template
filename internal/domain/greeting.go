package domain

import "time"

type Greeting struct {
	ID        string
	Message   string
	CreatedAt time.Time
}
