//go:generate mockgen -source=./parser.go -destination=./parser_mock.go -package=parser
package parser

import (
	"github.com/g3co/ticket-sorter/structs"
)

type IParser interface {
	Parse(card string) (*structs.Card, error)
}
