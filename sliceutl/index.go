package sliceutl

func InArray(key interface{}, sli []interface{}) bool {
	for _, v := range sli {
		if key == v {
			return true
		}
	}
	return false
}
