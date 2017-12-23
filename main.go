package main

import (
	"fmt"

	core "github.com/fredvdd/setgame/core"
)

func main() {
	cs := core.Cards()
	fmt.Println(len(cs), "cards")
	for i, c1 := range cs {
		for j, c2 := range cs {
			if i != j {
				var sc []core.Card
				for k, c3 := range cs {
					if k != i && k != j {
						if core.IsSet(c1, c2, c3) {
							sc = append(sc, c3)
						}
					}
				}
				if len(sc) != 1 {
					fmt.Println("unexpected set:", c1, c2, sc)
				} else {
					fmt.Println(c1, c2, sc[0])
				}
			}
		}
	}
}

/*
*   +---+ +---+ +---+
*   | ◈ | | S | | O |
*   | ◈ | | S | | O |
*   | ◈ | | S | | O |
*   +---+ +---+ +---+
 */
