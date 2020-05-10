// Package ticket_sorter provides sort operation for travel tickets
package ticket_sorter

import (
	"github.com/g3co/ticket-sorter/parser"
)

type TicketSorter struct {
	parser parser.IParser
}

type ITicketSorter interface {
	Sort(cards []string) ([]string, error)
}

// NewTicketSorter factory method for TicketSorter type
func NewTicketSorter(parser parser.IParser) TicketSorter {
	return TicketSorter{parser: parser}
}
