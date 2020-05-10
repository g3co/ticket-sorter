package ticket_sorter

import "github.com/g3co/ticket-sorter/structs"

func (a *TicketSorter) Sort(cards []string) (res []string, err error) {
	if len(cards) == 0 {
		return
	}

	res = make([]string, len(cards))

	fromRegistry := make(map[string]*structs.Card)
	toRegistry := make(map[string]*structs.Card)

	card := &structs.Card{}

	for _, item := range cards {
		card, err = a.parser.Parse(item)
		if err != nil {
			return res, err
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

	card = card.First()

	counter := 0
	for {
		res[counter] = card.Body
		if card = card.Next(); card == nil {
			break
		}
		counter++
	}

	if counter != len(cards)-1 {
		err = ErrInvalidCardSequence
		return
	}

	return
}
