package domain

import "errors"

type TicketType string

var (
	ErrTicketPriceZero = errors.New("Ticket price must be greater than zero")
)

const (
	TicketTypeHalf TicketType = "half" // Half-price ticket
	TicketTypeFull TicketType = "full" // Full-price ticket
)

type Ticket struct {
	ID         string
	EventID    string
	Spot       *Spot
	TicketType TicketType
	Price      float64
}

func isValidTicketType(ticketType TicketType) bool {
	return ticketType == TicketTypeHalf || ticketType == TicketTypeFull
}

func (t *Ticket) CalculatePrice() {
	if t.TicketType == TicketTypeHalf {
		t.Price /= 2
	}
	// } else {
	// 	t.Price = t.Spot.Event.Price
	// }
}

func (t *Ticket) Validate() error {
	if t.Price <= 0 {
		return ErrTicketPriceZero
	}
	return nil
}