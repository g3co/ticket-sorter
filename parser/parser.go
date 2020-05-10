//go:generate mockgen -source=./parser.go -destination=./mock/parser_mock.go
// Package parser provides parser for tickets for extract travel points information
package parser

import (
	"github.com/g3co/ticket-sorter/structs"
)

type IParser interface {
	Parse(card string) (*structs.Card, error)
}
