package connor

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEqual(t *testing.T) {
	Convey("$eq", t, func() {
		cases := []struct {
			conds string
			data  string

			match  bool
			hasErr bool
		}{
			{
				`{ "x": { "$eq": 1 } }`,
				`{ "x": 1, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "x": { "$eq": 2 } }`,
				`{ "x": 1, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "a": { "$eq": 1 } }`,
				`{ "x": 1, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "a.x": { "$eq": 1 } }`,
				`{ "a": { "x": 1 }, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "a": { "$eq": 1 } }`,
				`{ "a": { "x": 1 }, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "x": { "$eq": "1" } }`,
				`{ "x": 1, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "x": { "$eq": 1 } }`,
				`{ "x": "1", "y": 2 }`,
				false,
				false,
			},
			{
				`{ "x": { "z": 1 } }`,
				`{ "x": { "z": 1 }, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "x": null }`,
				`{ "x": null, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "x": 1 }`,
				`{ "x": null, "y": 2 }`,
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
