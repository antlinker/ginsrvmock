package json

// MarshalMustString json编码为字符串 不会返回错误,如果编码出错返回空字符串
func MarshalMustString(v interface{}) string {
	s, err := Marshal(v)
	if err != nil {
		return ""
	}
	return string(s)
}
