package ginsrvmock

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gogs.xiaoyuanjijiehao.com/tes/tesserver/internal/json"
)

// JSONAssert json断言器
type JSONAssert struct {
	Resp *httptest.ResponseRecorder
	t    *testing.T
}

// Assert 断言操作
// wantcode 想要获得的返回http状态码
// wantdata 想要获得的返回数据，首先wantdata会被json编码，然后在被解码为map[string]interface{}
// wanthead 只有第一个参数生效，用来比较返回的http头信息中是否包含　wanthead　包含的所有信息
func (a JSONAssert) Assert(wantcode int, wantdata interface{}, wanthead ...http.Header) {
	assert.Equal(a.t, wantcode, a.Resp.Code)
	if wantdata != nil {
		buff := a.Resp.Body.Bytes()
		tmp, err := json.Marshal(wantdata)
		assert.Empty(a.t, err)
		assert.JSONEq(a.t, string(tmp), string(buff))
	}

	if len(wanthead) == 0 || wanthead[0] == nil {
		return
	}
	for key := range wanthead[0] {
		wantheadvalue := wanthead[0].Get(key)
		assert.Equal(a.t, wantheadvalue, a.Resp.HeaderMap.Get(key))
	}
}
