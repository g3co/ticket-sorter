package ticket_sorter

func (a *TicketSort) buildChain(cards []string) (card *Card, err error) {
	card = &Card{}

	fromRegistry := make(map[string]*Card)
	toRegistry := make(map[string]*Card)

	for _, item := range cards {
		card, err = a.parseCard(item)
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
