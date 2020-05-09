package ticket_sorter

type Card struct {
	Body     string
	From     Location
	To       Location
	NextCard *Card
	PrevCard *Card
}

func (c *Card) FindFirst() (f *Card) {
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

type Location struct {
	Code  string
	Title string
}
