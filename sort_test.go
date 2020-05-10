package ticket_sorter

import (
	"errors"
	"github.com/g3co/ticket-sorter/parser"
	"github.com/g3co/ticket-sorter/structs"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestTicketSort_Sort(t *testing.T) {
	ctrl := gomock.NewController(t)
	ps := parser.NewMockParser(ctrl)

	ts := NewTicketSorter(ps)

	t.Run("succes", func(t *testing.T) {
		t.Run("zero card", func(t *testing.T) {
			card, err := ts.Sort([]string{})
			if err != nil {
				t.Error(err)
			}

			if len(card) != 0 {
				t.Error("card processed problem")
			}
		})

		t.Run("one card", func(t *testing.T) {
			cards := []string{"From [f:LA:LosAngeles] To [t:MIA:Miami]"}
			result := []string{"From LosAngeles To Miami"}

			ps.EXPECT().Parse(cards[0]).Return(&structs.Card{
				Body: result[0],
				From: struct {
					Code  string
					Title string
				}{Code: "LA", Title: "LosAngeles"},
				To: struct {
					Code  string
					Title string
				}{Code: "MIA", Title: "Miami"},
			}, nil)

			card, err := ts.Sort(cards)
			if err != nil {
				t.Error(err)
			}

			if card[0] != "From LosAngeles To Miami" {
				t.Error("card processed problem")
			}
		})

		t.Run("three card", func(t *testing.T) {
			input := []string{
				"From [f:LA:LosAngeles] To [t:MIA:Miami]",
				"From [f:WG:Washington] To [t:NY:NewYork]",
				"From [f:NY:NewYork] To [t:LA:LosAngeles]",
			}
			result := []string{
				"From Washington To NewYork",
				"From NewYork To LosAngeles",
				"From LosAngeles To Miami",
			}

			ps.EXPECT().Parse(input[0]).Return(&structs.Card{
				Body: result[2],
				From: struct {
					Code  string
					Title string
				}{Code: "LA", Title: "LosAngeles"},
				To: struct {
					Code  string
					Title string
				}{Code: "MIA", Title: "Miami"},
			}, nil).Times(1)

			ps.EXPECT().Parse(input[1]).Return(&structs.Card{
				Body: result[0],
				From: struct {
					Code  string
					Title string
				}{Code: "WG", Title: "Washington"},
				To: struct {
					Code  string
					Title string
				}{Code: "NY", Title: "NewYork"},
			}, nil).Times(1)

			ps.EXPECT().Parse(input[2]).Return(&structs.Card{
				Body: result[1],
				From: struct {
					Code  string
					Title string
				}{Code: "NY", Title: "NewYork"},
				To: struct {
					Code  string
					Title string
				}{Code: "LA", Title: "LosAngeles"},
			}, nil).Times(1)

			cards, err := ts.Sort(input)
			if err != nil {
				t.Error(err)
			}

			if cards[0] != result[0] {
				t.Error("card processed problem")
			}
		})
	})

	t.Run("errors", func(t *testing.T) {
		t.Run("parse err", func(t *testing.T) {
			cards := []string{"From [f:LA:LosAngeles] To [t:MIA:Miami]"}
			tErr := errors.New("parse err")

			ps.EXPECT().Parse(cards[0]).Return(nil, tErr)

			_, err := ts.Sort(cards)
			if err != tErr {
				t.Error("wrong error handler")
			}

		})

		t.Run("wrong chain card", func(t *testing.T) {
			input := []string{
				"From [f:LA:LosAngeles] To [t:MIA:Miami]",
				"From [f:WG:Washington] To [t:CH:Chicago]",
				"From [f:NY:NewYork] To [t:LA:LosAngeles]",
			}

			ps.EXPECT().Parse(input[0]).Return(&structs.Card{
				Body: "test",
				From: struct {
					Code  string
					Title string
				}{Code: "LA", Title: "LosAngeles"},
				To: struct {
					Code  string
					Title string
				}{Code: "MIA", Title: "Miami"},
			}, nil).Times(1)

			ps.EXPECT().Parse(input[1]).Return(&structs.Card{
				Body: "test",
				From: struct {
					Code  string
					Title string
				}{Code: "WG", Title: "Washington"},
				To: struct {
					Code  string
					Title string
				}{Code: "CH", Title: "Chicago"},
			}, nil).Times(1)

			ps.EXPECT().Parse(input[2]).Return(&structs.Card{
				Body: "test",
				From: struct {
					Code  string
					Title string
				}{Code: "NY", Title: "NewYork"},
				To: struct {
					Code  string
					Title string
				}{Code: "LA", Title: "LosAngeles"},
			}, nil).Times(1)

			_, err := ts.Sort(input)
			if err != ErrInvalidCardSequence {
				t.Error("wrong chain error handler is invalid")
			}
		})
	})
}
