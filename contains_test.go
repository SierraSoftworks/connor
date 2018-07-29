package connor

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestContains(t *testing.T) {
	Convey("Contains", t, func() {
		cases := []struct {
			conds string
			data  string

			match  bool
			hasErr bool
		}{
			{
				`{"x":{"$contains":"abc"}}`,
				`{"x":"abc"}`,
				true,
				false,
			},
			{
				`{"x":{"$contains":"bc"}}`,
				`{"x":"abc"}`,
				true,
				false,
			},
			{
				`{"x":{"$contains":"ab"}}`,
				`{"x":"abc"}`,
				true,
				false,
			},
			{
				`{"x":{"$contains":"xyz"}}`,
				`{"x":"abc"}`,
				false,
				false,
			},
			{
				`{"x":{"$contains":"abc"}}`,
				`{"x":1}`,
				false,
				true,
			},
		}

		for _, c := range cases {
			Convey(fmt.Sprintf("%s & %s", c.data, c.conds), func() {
				conds := prepConds(c.conds)
				data := prepData(c.data)

				m, err := Match(conds, data)
				if c.hasErr {
					So(err, ShouldNotBeNil)
				} else {
					So(err, ShouldBeNil)
				}

				So(m, ShouldEqual, c.match)
			})
		}
	})
}
