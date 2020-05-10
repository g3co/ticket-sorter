package ticket_sorter

import "github.com/g3co/ticket-sorter/structs"

func (a *TicketSort) buildChain(cards []string) (card *structs.Card, err error) {
	card = &structs.Card{}

	fromRegistry := make(map[string]*structs.Card)
	toRegistry := make(map[string]*structs.Card)

	for _, item := range cards {
		card, err = a.parser.Parse(item)
		if err != nil {
			return nil, err
		}

		fromRegistry[card.From.Code] = card
		if c, ok := toRegistry[card.From.Code]; ok {
			card.PrevCard = c
			c.NextCard = card
		}

		toRegistry[card.To.Code] = card
		if c, ok := fromRegistry[card.To.Code]; ok {
			card.NextCard = c
			c.PrevCard = card
		}
	}

	return
}
