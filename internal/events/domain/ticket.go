package domain

type TicketType string

const (
	TicketTypeHalf TicketType = "half" // Half-price ticket
	TicketTypeFull TicketType = "full" // Full-price ticket
)

type Ticket struct {
	ID       string
	EventID  string
	Spot     *Spot
	TicketID TicketType
	Price    float64
}