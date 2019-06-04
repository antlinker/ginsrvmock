package ginsrvmock

import (
	"net/http/httptest"
	"testing"
)

// ResultResp 返回请求结果
type ResultResp struct {
	Resp *httptest.ResponseRecorder
	t    *testing.T
}

// JSON 输出JSON断言器
func (r ResultResp) JSON() JSONAssert {
	return JSONAssert{
		Resp: r.Resp,
		t:    r.t,
	}
}
