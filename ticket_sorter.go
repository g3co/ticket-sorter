//go:generate mockgen -source=./ticket_sorter.go -destination=./ticket_sorter_mock.go -package=ticket_sorter
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

func NewTicketSorter(parser parser.IParser) TicketSorter {
	return TicketSorter{parser: parser}
}
