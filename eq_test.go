package connor

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEqual(t *testing.T) {
	Convey("$eq", t, func() {
		Convey("Complex Objects", func() {
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
				{
					`{ "x": { "$eq": { "y": 1, "z": 1 } } }`,
					`{ "x": { "y": 1, "z": 1 } }`,
					true,
					false,
				},
				{
					`{ "x": { "$eq": { "y": 2, "z": 2 } } }`,
					`{ "x": { "y": 1, "z": 1 } }`,
					false,
					false,
				},
				{
					`{ "x": { "$eq": { "y": [1] } } }`,
					`{ "x": { "y": [1], "z": 1 } }`,
					true,
					false,
				},
				{
					`{ "x": { "$eq": { "y": [2] } } }`,
					`{ "x": { "y": [1], "z": 1 } }`,
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

		Convey("Different Types", func() {
			cases := []struct {
				con  interface{}
				data interface{}

				match  bool
				hasErr bool
			}{
				{
					"test", "test",
					true, false,
				},
				{
					"test", 1,
					false, false,
				},
				{
					int8(10), 10,
					true, false,
				},
				{
					int16(10), 10,
					true, false,
				},
				{
					int32(10), 10,
					true, false,
				},
				{
					int64(10), 10,
					true, false,
				},
				{
					float32(10), 10,
					true, false,
				},
			}

			for _, c := range cases {
				conds := c.con
				data := c.data
				match := c.match
				hasErr := c.hasErr

				Convey(fmt.Sprintf("%T(%v) == %T(%v)", c.con, c.con, c.data, c.data), func() {
					m, err := Match(map[string]interface{}{
						"x": map[string]interface{}{"$eq": conds},
					}, map[string]interface{}{
						"x": data,
					})

					if hasErr {
						So(err, ShouldNotBeNil)
					} else {
						So(err, ShouldBeNil)
					}

					So(m, ShouldEqual, match)
				})
			}
		})

		Convey("Weird Combinations", func() {
			Convey(`{ "x": [1] } & { "x": [1] }`, func() {
				m, err := MatchWith("$eq", map[string]interface{}{
					"x": []int{1},
				}, map[string]interface{}{
					"x": []int{1},
				})

				So(err, ShouldBeNil)
				So(m, ShouldBeTrue)
			})

			Convey(`{ "x": [1, 2, 3] } & { "x": 1 }`, func() {
				m, err := MatchWith("$eq", map[string]interface{}{
					"x": 1,
				}, map[string]interface{}{
					"x": []interface{}{1, 2, 3},
				})

				So(err, ShouldBeNil)
				So(m, ShouldBeTrue)
			})

			Convey(`[{ "x": 1 }, { "x": 2 }, { "x": 3 }] & { "x": 1 }`, func() {
				m, err := MatchWith("$eq", map[string]interface{}{
					"x": 1,
				}, []interface{}{
					map[string]interface{}{"x": 1},
					map[string]interface{}{"x": 2},
					map[string]interface{}{"x": 3},
				})

				So(err, ShouldBeNil)
				So(m, ShouldBeTrue)
			})
		})
	})
}
