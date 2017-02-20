package connor

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLessEqual(t *testing.T) {
	Convey("$le", t, func() {
		cases := []struct {
			conds string
			data  string

			match  bool
			hasErr bool
		}{
			{
				`{ "x": { "$le": 2 } }`,
				`{ "x": 1, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "x": { "$le": 1 } }`,
				`{ "x": 1, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "x": { "$le": 0 } }`,
				`{ "x": 1, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "x": { "$le": 1.5 } }`,
				`{ "x": 1, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "a.x": { "$le": 2 } }`,
				`{ "a": { "x": 1 }, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "a": { "$le": 10 } }`,
				`{ "a": { "x": 1 }, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "x": { "$le": "0" } }`,
				`{ "x": 1, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "x": { "$le": 0 } }`,
				`{ "x": "1", "y": 2 }`,
				false,
				false,
			},
			{
				`{ "x": { "$le": "1" } }`,
				`{ "x": "1", "y": 2 }`,
				true,
				false,
			},
			{
				`{ "x": { "$le": "2" } }`,
				`{ "x": "1", "y": 2 }`,
				true,
				false,
			},
			{
				`{ "x": { "$le": "0" } }`,
				`{ "x": "1", "y": 2 }`,
				false,
				false,
			},
			{
				`{ "x": { "$le": [1] } }`,
				`{ "x": "1", "y": 2 }`,
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
