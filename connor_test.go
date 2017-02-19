package connor

import (
	"bytes"
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func prepConds(d string) map[string]interface{} {
	var v map[string]interface{}
	So(json.NewDecoder(bytes.NewBufferString(d)).Decode(&v), ShouldBeNil)
	return v
}

func prepData(d string) map[string]interface{} {
	var v map[string]interface{}
	So(json.NewDecoder(bytes.NewBufferString(d)).Decode(&v), ShouldBeNil)
	return v
}

func prepValue(d string) interface{} {
	var v interface{}
	So(json.NewDecoder(bytes.NewBufferString(d)).Decode(&v), ShouldBeNil)
	return v
}

func TestConnor(t *testing.T) {
	Convey("Connor", t, func() {
		Convey("Malformed Operator", func() {
			_, err := MatchWith("malformed", nil, nil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "operator should have '$' prefix")
		})

		Convey("Invalid Operator", func() {
			_, err := MatchWith("$invalid", nil, nil)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldEqual, "unknown operator 'invalid'")
		})
	})
}
