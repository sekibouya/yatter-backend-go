package object

import (
	"time"
)

type Status struct {
	ID        int       `json:"id,omitempty"`
	Account   Account   `json:"account,omitempty"`
	URL       *string   `json:"url,omitempty" db:"url"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}

func NewStatus(account Account, content string) *Status {
	return &Status{
		Account:   account,
		Content:   content,
		CreatedAt: time.Now(),
	}
}
