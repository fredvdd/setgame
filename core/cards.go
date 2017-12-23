package core

// Foo is bar
const Foo = 3

// Number is the number of shapes on a Card
type Number int

// Symbol is the name of a shape on a Card
type Symbol string

// Shading is the pattern of the shapes on a Card
type Shading string

// Color is color of the shapes on a Card
type Color string

const (
	// One is the Number for 1 shape on a Card
	One Number = iota + 1
	// Two is the Number for 2 shapes on a Card
	Two
	// Three is the Number for 3 shapes on a Card
	Three
)

const (
	// Diamond is the diamond shape
	Diamond Symbol = "diamond"
	// Squiggle is the squiggle shape
	Squiggle Symbol = "squiggle"
	// Oval is oval shape
	Oval Symbol = "oval"
)

const (
	// Solid is the solid shading
	Solid Shading = "solid"
	// Striped is the striped shading
	Striped Shading = "striped"
	// Open is the open shading
	Open Shading = "open"
)

const (
	// Red is the red color
	Red Color = "red"
	// Green is the green color
	Green Color = "green"
	// Purple is the purple color
	Purple Color = "purple"
)

// Card represents a card in the Set game
type Card struct {
	Number  Number
	Symbol  Symbol
	Shading Shading
	Color   Color
}

var cards []Card

func init() {
	cs := []Card{}
	for _, num := range []Number{One, Two, Three} {
		for _, sym := range []Symbol{Diamond, Squiggle, Oval} {
			for _, sha := range []Shading{Solid, Striped, Open} {
				for _, col := range []Color{Red, Green, Purple} {
					cs = append(cs, Card{num, sym, sha, col})
				}
			}
		}
	}
	cards = cs
}

// Cards returns the full deck of Set cards
func Cards() []Card {
	return cards
}

// IsSet determines if the provided cards form a set
func IsSet(a, b, c Card) bool {
	if a == b && b == c {
		return false
	}
	return isPropertySet(a.Number, b.Number, c.Number) &&
		isPropertySet(a.Symbol, b.Symbol, c.Symbol) &&
		isPropertySet(a.Shading, b.Shading, c.Shading) &&
		isPropertySet(a.Color, b.Color, c.Color)
}

func isPropertySet(a, b, c interface{}) bool {
	return (a == b && b == c) ||
		(a != b && b != c && a != c)
}
