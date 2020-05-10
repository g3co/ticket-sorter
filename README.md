# Travel Card Sorter

Library for sorting travel tickets by destination points

### Installation

Add to your go.mod file
`require github.com/g3co/ticket_sorter latest`

### Test


`go test . ./parser`

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
        "From [f:C:PointC] To [t:D:PointD] by bus, place 12",
        "Airport [f:D:PointD] bus stop [t:F:PointF] Gate 45B. Seat 3A.\nBaggage drop at ticket counter 344",
        "From [f:A:PointА] To [t:B:PointB] No seat assignment.",
        "From [f:F:PointF] To [t:E:PointE] Seat 13A.",
        "Downtown [f:B:PointB] by car [t:C:PointC] Terminal 22, Seat 3A.",
    }

	ts := ticketSorter.NewTicketSorter(parser.NewCardParser())

	cards, err := ts.Sort(sample)
	if err != nil {
		panic(err)
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
Destination f(from) or t(to), unique place name A and B, 
and name for replacement PointА and PointB, separated by ":" key. 

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