package connor

import (
	"strconv"
	"strings"
)

func getField(data map[string]interface{}, field string) interface{} {
	fps := strings.Split(field, ".")
	d := interface{}(data)
	for _, fp := range fps {
		switch td := d.(type) {
		case map[string]interface{}:
			f, ok := td[fp]
			if !ok {
				return nil
			}

			d = f
		case []interface{}:
			fpi, err := strconv.Atoi(fp)
			if err != nil {
				return nil
			}
			d = td[fpi]
		default:
			return nil
		}
	}

	return d
}
