package parser

import (
	"testing"
)

func TestCardParser_Parse(t *testing.T) {
	cp := NewCardParser()

	t.Run("succes", func(t *testing.T) {
		card, err := cp.Parse("From [f:LA:LosAngeles] To [t:MIA:Miami]")
		if err != nil {
			t.Error(err)
		}

		if card.Body != "From LosAngeles To Miami" {
			t.Error("wrong card body")
		}
	})

	t.Run("wrong format", func(t *testing.T) {
		_, err := cp.Parse("From f:LA:LosAngeles To [t:MIA:Miami]")
		if err == nil {
			t.Error("passed wrong format")
		}
	})

	t.Run("wrong direction", func(t *testing.T) {
		_, err := cp.Parse("From [x:LA:LosAngeles] To [t:MIA:Miami]")
		if err == nil {
			t.Error("passed wrong direction format")
		}
	})
}
