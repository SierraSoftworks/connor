package connor

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEqualsSensu(t *testing.T) {
	Convey("$eq", t, func() {

		Convey("Sensu Check", func() {
			conds := prepConds(`{
				"check.status": 0
			}`)
			data := map[string]interface{}{
				"client": "razz-base-stackstorm",
				"check": map[string]interface{}{
					"command":     "/opt/sensu/embedded/bin/check-cpu.rb",
					"handlers":    []interface{}{"default", "st2"},
					"name":        "CPU",
					"issued":      1487492532,
					"interval":    60,
					"subscribers": []interface{}{"generic"},
					"executed":    1487492532,
					"duration":    1.115,
					"output":      "This is a quick test",
					"status":      0,
					"remediations": []interface{}{
						map[string]interface{}{
							"name":    "all good",
							"command": "echo 'OK' > check_handler.dat",
							"conditions": map[string]interface{}{
								"check.status": 0,
							},
						},
						map[string]interface{}{
							"name":    "so-so",
							"command": "echo 'WARN' > check_handler.dat",
							"conditions": map[string]interface{}{
								"check.status": 1,
								"occurrences":  2,
							},
						},
						map[string]interface{}{
							"name":    "it's on fire!",
							"command": "echo 'CRIT' > check_handler.dat",
							"conditions": map[string]interface{}{
								"check.status": 2,
								"occurrences":  2,
							},
						},
					},
				},
			}

			m, err := Match(conds, data)
			So(err, ShouldBeNil)
			So(m, ShouldBeTrue)
		})
	})
}
