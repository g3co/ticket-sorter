package ticket_sorter

import (
	"regexp"
)

type TicketSort struct {
	Matcher *regexp.Regexp
}

func TicketSorter(cards []string) ([]string, error) {
	matcher := regexp.MustCompile(`\[([f|t]):([^:]+):([^\]]+)\]`)
	app := TicketSort{Matcher: matcher}
	return app.process(cards)
}

func (a *TicketSort) process(cards []string) (res []string, err error) {
	c, err := a.buildChain(cards)
	if err != nil {
		return
	}

	res = make([]string, len(cards))
	counter := 0
	for {
		res[counter] = c.Body
		if c.NextCard == nil {
			break
		}

		c = c.NextCard
		counter++
	}

	return res, nil
}
