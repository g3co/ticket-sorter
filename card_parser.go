package ticket_sorter

import (
	"errors"
	"strings"
)

const (
	LocationKeyFrom        = "f"
	LocationKeyTo          = "t"
	LocationIndexPattern   = 0
	LocationIndexDirection = 1
	LocationIndexCode      = 2
	LocationIndexTitle     = 3
)

var (
	ErrWrongCardFormat = errors.New("wrong card format")
)

func (a *TicketSort) parseCard(card string) (c *Card, err error) {

	result := a.Matcher.FindAllStringSubmatch(card, 2)
	if len(result) != 2 {
		err = ErrWrongCardFormat
		return
	}

	c = &Card{}

	for _, item := range result {
		l := Location{
			Code:  item[LocationIndexCode],
			Title: item[LocationIndexTitle],
		}

		if item[LocationIndexDirection] == LocationKeyFrom {
			c.From = l
		} else if item[LocationIndexDirection] == LocationKeyTo {
			c.To = l
		} else {
			err = ErrWrongCardFormat
			return
		}

		card = strings.Replace(card, item[LocationIndexPattern], l.Title, 1)
	}

	c.Body = card

	return
}
