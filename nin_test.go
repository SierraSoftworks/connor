package connor

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNotIn(t *testing.T) {
	Convey("$nin", t, func() {
		cases := []struct {
			conds string
			data  string

			match  bool
			hasErr bool
		}{
			{
				`{ "x": { "$nin": [1] } }`,
				`{ "x": 1, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "x": { "$nin": [2] } }`,
				`{ "x": 1, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "x": { "$nin": [1, 2, 3] } }`,
				`{ "x": 1, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "a.x": { "$nin": [1] } }`,
				`{ "a": { "x": 1 }, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "a": { "$nin": [1] } }`,
				`{ "a": { "x": 1 }, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "a": { "$nin": 1 } }`,
				`{ "a": { "x": 1 }, "y": 2 }`,
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
