package ticket_sorter

type Card struct {
	Body     string
	From     Location
	To       Location
	NextCard *Card
	PrevCard *Card
}

func (c *Card) First() (f *Card) {
	f = c

	if f.PrevCard == nil {
		return
	}

	for {
		f = f.PrevCard
		if f.PrevCard == nil {
			break
		}
	}

	return
}

func (c *Card) Next() *Card {
	return c.NextCard
}

type Location struct {
	Code  string
	Title string
}
