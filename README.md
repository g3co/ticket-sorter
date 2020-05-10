# Travel Card Sorter

Library for sorting travel tickets by destination point

### Installation

Add to your go.mod file
`require github.com/g3co/ticket_sorter latest`

### Usage



### Sample application
```
package main

import (
	"fmt"
	"github.com/g3co/ticket_sorter"
	"github.com/g3co/ticket_sorter/parser"
)

func main() {
	sample := []string{
		"From [f:A:PointА] To [t:B:PointB]",
		"From [f:C:PointC] To [t:D:PointD]",
		"From [f:B:PointB] To [t:C:PointC]",
		"From [f:F:PointF] To [t:E:PointE]",
		"From [f:D:PointD] To [t:F:PointF]",
	}

	ts := ticket_sorter.NewTicketSorter(parser.NewCardParser())

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

