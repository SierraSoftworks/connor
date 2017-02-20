package connor

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLess(t *testing.T) {
	Convey("$lt", t, func() {
		cases := []struct {
			conds string
			data  string

			match  bool
			hasErr bool
		}{
			{
				`{ "x": { "$lt": 2 } }`,
				`{ "x": 1, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "x": { "$lt": 1 } }`,
				`{ "x": 1, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "x": { "$lt": 1.5 } }`,
				`{ "x": 1, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "a.x": { "$lt": 2 } }`,
				`{ "a": { "x": 1 }, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "a": { "$lt": 10 } }`,
				`{ "a": { "x": 1 }, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "x": { "$lt": "0" } }`,
				`{ "x": 1, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "x": { "$lt": 0 } }`,
				`{ "x": "1", "y": 2 }`,
				false,
				false,
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
