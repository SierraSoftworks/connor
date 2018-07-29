package connor

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumbersHelpers(t *testing.T) {
	Convey("Numbers Helpers", t, func() {
		Convey("tryUpcastNumber", func() {
			cases := []struct {
				in  interface{}
				out interface{}
			}{
				{int8(10), int64(10)},
				{int16(10), int64(10)},
				{int32(10), int64(10)},
				{int64(10), int64(10)},
				{float32(10), float64(10)},
				{float64(10), float64(10)},
				{"test", "test"},
			}

			for _, c := range cases {
				in := c.in
				out := c.out

				Convey(fmt.Sprintf("%T(%v) -> %T(%v)", in, in, out, out), func() {
					So(tryUpcastNumber(in), ShouldEqual, out)
					So(tryUpcastNumber(in), ShouldHaveSameTypeAs, out)
				})
			}
		})

		Convey("numbersEqual", func() {
			cases := []struct {
				a interface{}
				b interface{}

				matches bool
			}{
				{int64(10), int64(10), true},
				{int64(10), int64(0), false},

				{float64(10), int64(10), true},
				{int64(10), float64(10), true},
				{float64(10), int64(12), false},
				{int64(10), float64(12), false},

				{int8(5), float64(5), true},

				{"test", int64(5), false},
				{int64(5), "test", false},
				{"test", float64(5), false},
				{float64(5), "test", false},
				{"test", "test", false},
			}

			for _, c := range cases {
				a := c.a
				b := c.b
				expected := c.matches
				Convey(fmt.Sprintf("%T(%v) == %T(%v)", a, a, b, b), func() {
					if expected {
						So(numbersEqual(a, b), ShouldBeTrue)
					} else {
						So(numbersEqual(a, b), ShouldBeFalse)
					}
				})
			}
		})
	})
}
