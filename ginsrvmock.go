package ginsrvmock

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"gogs.xiaoyuanjijiehao.com/tes/tesserver/internal/json"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gogs.xiaoyuanjijiehao.com/tes/tesserver/internal/web/ginsrv"
)

// GinMock gin mock接口
type GinMock interface {
	Bind(t *testing.T) GinMock
	Header(head http.Header) GinMock
	Get(url string) ResultResp
	Delete(url string) ResultResp
	Put(url string, body interface{}) ResultResp
	Post(url string, body interface{}) ResultResp
	Requst(req *http.Request) ResultResp
	// 矩阵测试
	Matrix() Matrix
}

// New 创建gin服务mock
func New(handler func(route *ginsrv.Engine)) GinMock {
	r := gin.Default()
	if handler != nil {
		handler(r)
	}
	return ginMock{
		route: r,
	}
}

type ginMock struct {
	route *ginsrv.Engine
	t     *testing.T
	head  http.Header
}

func (g ginMock) Bind(t *testing.T) GinMock {
	g.t = t
	return g
}
func (g ginMock) Header(head http.Header) GinMock {
	g.head = head
	return g
}
func (g ginMock) Matrix() Matrix {
	return Matrix{g}
}
func (g ginMock) Get(url string) ResultResp {
	req, err := http.NewRequest("GET", url, nil)
	assert.Empty(g.t, err)
	return g.Requst(req)
}

func (g ginMock) Delete(url string) ResultResp {
	req, err := http.NewRequest("DELETE", url, nil)
	assert.Empty(g.t, err)
	return g.Requst(req)
}
func (g ginMock) Put(url string, body interface{}) ResultResp {
	buffer := bytes.NewBuffer(nil)
	err := json.NewEncoder(buffer).Encode(body)
	assert.Empty(g.t, err)
	req, err := http.NewRequest("PUT", url, buffer)
	assert.Empty(g.t, err)
	return g.Requst(req)
}
func (g ginMock) Post(url string, body interface{}) ResultResp {
	buffer := bytes.NewBuffer(nil)
	err := json.NewEncoder(buffer).Encode(body)
	assert.Empty(g.t, err)
	req, err := http.NewRequest("POST", url, buffer)
	assert.Empty(g.t, err)
	return g.Requst(req)
}

func (g ginMock) Requst(req *http.Request) ResultResp {
	w := httptest.NewRecorder()
	g.route.ServeHTTP(w, req)
	return ResultResp{
		Resp: w,
		t:    g.t,
	}
}
