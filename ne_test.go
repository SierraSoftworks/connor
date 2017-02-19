package connor

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNotEqual(t *testing.T) {
	Convey("$ne", t, func() {
		cases := []struct {
			conds string
			data  string

			match  bool
			hasErr bool
		}{
			{
				`{ "x": { "$ne": 1 } }`,
				`{ "x": 1, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "x": { "$ne": 2 } }`,
				`{ "x": 1, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "a": { "$ne": 1 } }`,
				`{ "x": 1, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "a.x": { "$ne": 1 } }`,
				`{ "a": { "x": 1 }, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "a": { "$ne": 1 } }`,
				`{ "a": { "x": 1 }, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "x": { "$ne": "1" } }`,
				`{ "x": 1, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "x": { "$ne": 1 } }`,
				`{ "x": "1", "y": 2 }`,
				true,
				false,
			},
			{
				`{ "x": { "$ne": { "z": 1 } } }`,
				`{ "x": { "z": 1 }, "y": 2 }`,
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
