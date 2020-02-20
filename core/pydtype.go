package core

import "gtpython/number"

//表示python的数据类型
type PyValue interface{}

func Add(val1 PyValue, val2 PyValue) int64 {
	val1_int, _ := convertToInteger(val1)
	val2_int, _ := convertToInteger(val2)
	res := val1_int + val2_int
	return res

}

func typeOf(val PyValue) PyType {
	switch val.(type) {
	case nil:
		return PY_NONE //go里面的nil 对应 python 的none
	case bool:
		return PY_Bool
	case int64:
		return PY_INT
	case float64:
		return PY_Float
	case string:
		return PY_STRING
	case complex64:
		return PY_COMPLEX
	default:
		panic("type error")

	}

}

func convertToBoolean(val PyValue) bool {
	switch x := val.(type) {
	case nil:
		return false
	case bool:
		return x
	default:
		return true
	}
}
func convertToFloat(val PyValue) (float64, bool) {
	switch x := val.(type) {
	case float64:
		return x, true
	case int64:
		return float64(x), true
	case string:
		return number.ParseFloat(x)
	default:
		return 0, false
	}
}
func convertToInteger(val PyValue) (int64, bool) {
	switch x := val.(type) {
	case int:
		return int64(x), true
	case int64:
		return x, true
	case float64:
		return number.FloatToInteger(x)
	case string:
		return _stringToInteger(x)
	default:
		return 0, false
	}
}

func _stringToInteger(s string) (int64, bool) {
	if i, ok := number.ParseInteger(s); ok {
		return i, true
	}
	if f, ok := number.ParseFloat(s); ok {
		return number.FloatToInteger(f)
	}
	return 0, false
}
