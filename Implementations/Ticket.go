package Implementations

import (
	"github.com/google/uuid"
)

type Ticket struct {
	ticketID string
}

func NewTicket() *Ticket {
	return &Ticket{
		ticketID: uuid.NewString(),
	}
}

func (t *Ticket) Equals(other *Ticket) bool {
	return t.ticketID == other.ticketID
}
