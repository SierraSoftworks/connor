package connor

import (
	"reflect"
	"strings"
)

func init() {
	Register(&EqualOperator{})
}

// EqualOperator is an operator which performs object equality
// tests.
type EqualOperator struct {
}

func (o *EqualOperator) Name() string {
	return "eq"
}

func (o *EqualOperator) Evaluate(condition, data interface{}) (bool, error) {
	switch cn := condition.(type) {
	case string:
		if d, ok := data.(string); ok {
			return d == cn, nil
		}
		return false, nil
	case int:
		if d, ok := data.(int); ok {
			return d == cn, nil
		}
		return false, nil
	case float32:
		if d, ok := data.(float32); ok {
			return d == cn, nil
		}
		return false, nil
	case float64:
		if d, ok := data.(float64); ok {
			return d == cn, nil
		}
		return false, nil
	case map[string]interface{}:
		m := true
		for prop, cond := range cn {
			if !m {
				// No need to evaluate after we fail
				continue
			}

			if strings.HasPrefix(prop, "$") {
				mm, err := MatchWith(prop, cond, data)
				if err != nil {
					return false, err
				}

				m = m && mm
			} else if d, ok := data.(map[string]interface{}); ok {
				mm, err := MatchWith("$eq", cond, getField(d, prop))
				if err != nil {
					return false, err
				}

				m = m && mm
			} else {
				return reflect.DeepEqual(condition, data), nil
			}
		}

		return m, nil
	default:
		return reflect.DeepEqual(condition, data), nil
	}
}
