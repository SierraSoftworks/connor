package connor

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFields(t *testing.T) {
	Convey("Fields", t, func() {
		cases := []struct {
			data  string
			field string
			value string
		}{
			{
				`{"x": 1}`,
				"x",
				"1",
			},
			{
				`{"x": 1}`,
				"y",
				"null",
			},
			{
				`{"x": { "y": 1 }}`,
				"x.y",
				"1",
			},
			{
				`{"x": { "y": 1 }}`,
				"x",
				`{ "y": 1 }`,
			},
			{
				`{"x": 1}`,
				"x.y",
				`null`,
			},
			{
				`{"x": [ { "y": 1 }, { "y" : 2 } ]}`,
				"x.0.y",
				`1`,
			},
			{
				`{"x": [ { "y": 1 }, { "y" : 2 } ]}`,
				"x.1.y",
				`2`,
			},
			{
				`{"x": [ { "y": [ 5,6] }]}`,
				"x.0.y.0",
				`5`,
			},
			{
				`{"x": [ { "y": [ 5,6] }]}`,
				"x.0.y.1",
				`6`,
			},
			{
				`{"x": [ { "y": [ 5,6] }]}`,
				"x.0.y.2",
				`null`,
			},
			{
				`{"x": [ { "y": [ 5,6] }]}`,
				"x.-1.y.2",
				`null`,
			},
			{
				`{"x": [ { "y": 1 } ]}`,
				"x.0.z",
				`null`,
			},
			{
				`{"x": [ { "y": 1 } ]}`,
				"x.3.z",
				`null`,
			},
		}

		for _, c := range cases {
			Convey(fmt.Sprintf("%s | %s == %v", c.field, c.data, c.value), func() {
				d := prepData(c.data)
				ev := prepValue(c.value)
				v := getField(d, c.field)
				So(v, ShouldResemble, ev)
			})
		}
	})
}
