//go:generate mockgen -source=./ticket_sorter.go -destination=./ticket_sorter_mock.go -package=ticket_sorter
package ticket_sorter

import (
	"github.com/g3co/ticket-sorter/parser"
)

type TicketSort struct {
	parser parser.IParser
}

type ITicketSort interface {
	Sort(cards []string) ([]string, error)
}

func NewTicketSorter(parser parser.IParser) TicketSort {
	return TicketSort{parser: parser}
}
