package core

import (
	"testing"
)

func TestIsSet(t *testing.T) {
	tests := []struct {
		desc    string
		a, b, c Card
		want    bool
	}{
		{desc: "set, color different",
			a:    Card{One, Oval, Open, Red},
			b:    Card{One, Oval, Open, Green},
			c:    Card{One, Oval, Open, Purple},
			want: true,
		},
		{desc: "set, shading different",
			a:    Card{One, Oval, Open, Red},
			b:    Card{One, Oval, Solid, Red},
			c:    Card{One, Oval, Striped, Red},
			want: true,
		},
		{desc: "set, shape different",
			a:    Card{One, Oval, Open, Red},
			b:    Card{One, Squiggle, Open, Red},
			c:    Card{One, Diamond, Open, Red},
			want: true,
		},
		{desc: "set, number different",
			a:    Card{One, Oval, Open, Red},
			b:    Card{Two, Oval, Open, Red},
			c:    Card{Three, Oval, Open, Red},
			want: true,
		},
		{desc: "not set, identical (impossible)",
			a:    Card{One, Oval, Open, Red},
			b:    Card{One, Oval, Open, Red},
			c:    Card{One, Oval, Open, Red},
			want: false,
		},
		{desc: "not set, number split",
			a:    Card{One, Oval, Open, Red},
			b:    Card{Two, Oval, Open, Red},
			c:    Card{Two, Oval, Open, Red},
			want: false,
		},
		{desc: "not set, shape split",
			a:    Card{One, Oval, Open, Red},
			b:    Card{One, Squiggle, Open, Red},
			c:    Card{One, Squiggle, Open, Red},
			want: false,
		},
		{desc: "not set, shading split",
			a:    Card{One, Oval, Open, Red},
			b:    Card{One, Oval, Striped, Red},
			c:    Card{One, Oval, Striped, Red},
			want: false,
		},
		{desc: "not set, color split",
			a:    Card{One, Oval, Open, Red},
			b:    Card{One, Oval, Open, Red},
			c:    Card{One, Oval, Open, Purple},
			want: false,
		},
		{desc: "set, only color same",
			a:    Card{One, Oval, Open, Red},
			b:    Card{Two, Squiggle, Striped, Red},
			c:    Card{Three, Diamond, Solid, Red},
			want: true,
		},
	}
	for _, test := range tests {
		got := IsSet(test.a, test.b, test.c)
		if test.want != got {
			t.Errorf("test IsSet(%q), wanted %t, got %t", test.desc, test.want, got)
		}
	}
}
