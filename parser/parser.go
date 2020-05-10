package parser

import (
	"github.com/g3co/ticket_sorter/structs"
)

type Parser interface {
	Parse(card string) (*structs.Card, error)
}
