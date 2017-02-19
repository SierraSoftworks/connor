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
