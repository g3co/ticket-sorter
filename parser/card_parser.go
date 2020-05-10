package parser

import (
	"errors"
	"github.com/g3co/ticket-sorter/structs"
	"regexp"
	"strings"
)

const (
	pointPattern = `\[([f|t]):([^:]+):([^\]]+)\]`

	LocationKeyFrom        = "f"
	LocationKeyTo          = "t"
	LocationIndexPattern   = 0
	LocationIndexDirection = 1
	LocationIndexCode      = 2
	LocationIndexTitle     = 3
)

type CardParser struct {
	matcher *regexp.Regexp
}

var (
	ErrWrongCardFormat = errors.New("wrong card format")
)

func NewCardParser() *CardParser {
	matcher := regexp.MustCompile(pointPattern)
	return &CardParser{matcher: matcher}
}

func (cp *CardParser) Parse(card string) (c *structs.Card, err error) {

	result := cp.matcher.FindAllStringSubmatch(card, 2)
	if len(result) != 2 {
		err = ErrWrongCardFormat
		return
	}

	c = &structs.Card{}

	for _, item := range result {
		l := structs.Location{
			Code:  item[LocationIndexCode],
			Title: item[LocationIndexTitle],
		}

		if item[LocationIndexDirection] == LocationKeyFrom {
			c.From = l
		} else if item[LocationIndexDirection] == LocationKeyTo {
			c.To = l
		}

		card = strings.Replace(card, item[LocationIndexPattern], l.Title, 1)
	}

	c.Body = card

	return
}
