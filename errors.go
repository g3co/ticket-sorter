package ticket_sorter

import "errors"

var (
	ErrInvalidCardSequence = errors.New("invalid card sequence")
	ErrWrongCardFormat     = errors.New("wrong card format")
)
