package ticket_sorter

func (a *TicketSort) Sort(cards []string) (res []string, err error) {
	if len(cards) == 0 {
		return
	}

	c, err := a.buildChain(cards)
	if err != nil {
		return
	}

	c = c.First()

	res = make([]string, len(cards))

	counter := 0
	for {
		res[counter] = c.Body
		if c = c.Next(); c == nil {
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
