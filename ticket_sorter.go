package ticket_sorter

import (
	"github.com/g3co/ticket-sorter/parser"
)

type TicketSort struct {
	parser parser.Parser
}

type ITicketSort interface {
	Sort(cards []string) ([]string, error)
}

func NewTicketSorter(parser parser.Parser) TicketSort {
	return TicketSort{parser: parser}
}
