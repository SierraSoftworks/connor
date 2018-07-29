package connor

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLess(t *testing.T) {
	now := time.Now()

	Convey("$lt", t, func() {
		Convey("Basic Cases", func() {
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

		Convey("Different Types", func() {
			cases := []struct {
				con  interface{}
				data interface{}

				match  bool
				hasErr bool
			}{
				{
					"abc", "def",
					false, false,
				},
				{
					"abc", "abc",
					false, false,
				},
				{
					"abc", "aaa",
					true, false,
				},
				{
					"test", 1,
					false, false,
				},
				{
					int8(10), 5,
					true, false,
				},
				{
					int16(10), 1,
					true, false,
				},
				{
					int32(10), 3,
					true, false,
				},
				{
					int64(10), 9,
					true, false,
				},
				{
					int64(10), 10,
					false, false,
				},
				{
					int64(10), 12,
					false, false,
				},
				{
					float32(10), 9,
					true, false,
				},
				{
					int64(10), float32(9),
					true, false,
				},
				{
					int64(10), "test",
					false, false,
				},
				{
					now, now,
					false, false,
				},
				{
					now, now.Add(time.Second),
					false, false,
				},
				{
					now, now.Add(-time.Second),
					true, false,
				},
				{
					now, 10,
					false, false,
				},
				{
					[]int{10}, []int{12},
					false, true,
				},
			}

			for _, c := range cases {
				conds := c.con
				data := c.data
				match := c.match
				hasErr := c.hasErr

				Convey(fmt.Sprintf("%T(%v) == %T(%v)", c.con, c.con, c.data, c.data), func() {
					m, err := Match(map[string]interface{}{
						"x": map[string]interface{}{"$lt": conds},
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
	})
}
