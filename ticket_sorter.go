package ticket_sorter

import (
	"regexp"
)

const placePattern = `\[([f|t]):([^:]+):([^\]]+)\]`

type TicketSort struct {
	matcher *regexp.Regexp
}

type ITicketSort interface {
	Process(cards []string) ([]string, error)
}

func NewTicketSorter() TicketSort {
	matcher := regexp.MustCompile(placePattern)
	return TicketSort{matcher: matcher}
}
