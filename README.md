# Travel Card Sorter

Library for sorting travel tickets by destination points

### Installation

Add to your go.mod file
`require github.com/g3co/ticket_sorter latest`

### Usage

Create a new instance of TicketSorter and inject a parser for the card.

```ticketSorter.NewTicketSorter(parser.NewCardParser())```

Also, you can implement a self-made parser for your card types.

### Sample application
```
package main

import (
	"fmt"
	ticketSorter "github.com/g3co/ticket-sorter"
	"github.com/g3co/ticket-sorter/parser"
)

func main() {
	sample := []string{
		"From [f:A:PointА] To [t:B:PointB]",
		"From [f:C:PointC] To [t:D:PointD]",
		"From [f:B:PointB] To [t:C:PointC]",
		"From [f:F:PointF] To [t:E:PointE]",
		"From [f:D:PointD] To [t:F:PointF]",
	}

	ts := ticketSorter.NewTicketSorter(parser.NewCardParser())

	cards, err := ts.Process(sample)
	if err != nil {
		fmt.Println(err.Error())
		panic("")
	}

	for _, card := range cards {
		fmt.Println(" * ", card)
	}
}
```

### Input data format for standard parser
The library works with array of string. Each card is a text string with point anchors.
Each card must contain two anchors for the start and end points.

`[f:A:PointА] [t:B:PointB]`

Each anchor contains 3 entities. 
Destination f(from) or t(to), unique name A and B for destination, 
and name for replacement, separated by ":" key. 

For example:
```
[]string{
    "From [f:LA:Los Angeles airport] To [t:MIA:Miami]",
    "From [f:NY:New-York] To [t:LA:Los Angeles]",
    "From [f:MIA:Miami Beach bus terminal] To [t:KW:Key West]",
}
```

### Output data format
```
[]string{
    "From New-York To Los Angeles",
    "From Los Angeles airport To Miami",
    "From Miami Beach bus terminal To Key West",
}
```