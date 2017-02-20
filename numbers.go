package connor

func numbersEqual(condition, data interface{}) bool {
	uc := tryUpcastNumber(condition)
	ud := tryUpcastNumber(data)

	switch ucv := uc.(type) {
	case int64:
		if udv, ok := ud.(int64); ok {
			return ucv == udv
		} else if udv, ok := ud.(float64); ok {
			fucv := float64(ucv)
			return fucv == udv && int64(fucv) == ucv
		}
		return false
	case float64:
		if udv, ok := ud.(float64); ok {
			return ucv == udv
		} else if udv, ok := ud.(int64); ok {
			iucv := int64(ucv)
			return iucv == udv && float64(iucv) == ucv
		}
		return false
	default:
		return false
	}
}

func tryUpcastNumber(n interface{}) interface{} {
	switch nn := n.(type) {
	case int8:
		return int64(nn)
	case int16:
		return int64(nn)
	case int32:
		return int64(nn)
	case int:
		return int64(nn)
	case int64:
		return nn
	case float32:
		return float64(nn)
	case float64:
		return nn
	default:
		return n
	}
}
