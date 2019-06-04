package ginsrvmock

import (
	"net/http"
	"testing"
)

// RespType 响应类型
type RespType int

const (
	// RespTypeJSON json类型
	RespTypeJSON RespType = 1
)

// MatrixTestArg 测试参数
type MatrixTestArg struct {
	Name     string
	URL      string
	Body     interface{}
	Head     http.Header
	Method   string
	WantCode int
	WantData interface{}
	WantHead http.Header
	RespType RespType
}

// Matrix 矩阵测试
type Matrix struct {
	ginMock
}

// NewGetArg 创建get参数
func NewGetArg(name, url string, wantCode int, wantData interface{}) MatrixTestArg {
	return MatrixTestArg{
		Name:     name,
		URL:      url,
		Method:   http.MethodGet,
		WantCode: wantCode,
		WantData: wantData,
		RespType: RespTypeJSON,
	}

}

// NewPutArg 创建put参数
func NewPutArg(name, url string, body interface{}, wantCode int, wantData interface{}) MatrixTestArg {
	return MatrixTestArg{
		Name:     name,
		URL:      url,
		Method:   http.MethodPut,
		Body:     body,
		WantCode: wantCode,
		WantData: wantData,
		RespType: RespTypeJSON,
	}

}

// NewPostArg 创建post参数
func NewPostArg(name, url string, body interface{}, wantCode int, wantData interface{}) MatrixTestArg {
	return MatrixTestArg{
		Name:     name,
		URL:      url,
		Method:   http.MethodPost,
		Body:     body,
		WantCode: wantCode,
		WantData: wantData,
		RespType: RespTypeJSON,
	}

}

// NewDeleteArg 创建delete参数
func NewDeleteArg(name, url string, wantCode int, wantData interface{}) MatrixTestArg {
	return MatrixTestArg{
		Name:     name,
		URL:      url,
		Method:   http.MethodDelete,
		WantCode: wantCode,
		WantData: wantData,
		RespType: RespTypeJSON,
	}

}

func (g Matrix) wantAssert(resp ResultResp, arg MatrixTestArg) {
	switch arg.RespType {
	case RespTypeJSON:
		resp.JSON().Assert(arg.WantCode, arg.WantData, arg.WantHead)
	default:
		g.t.Errorf("暂时不支持返回的编码格式")
	}
}
func (g Matrix) getTest(arg MatrixTestArg) {
	g.t.Run(arg.Name, func(t *testing.T) {
		g.wantAssert(g.Bind(t).Header(arg.Head).Get(arg.URL), arg)
	})
}
func (g Matrix) deleteTest(arg MatrixTestArg) {
	g.t.Run(arg.Name, func(t *testing.T) {
		g.wantAssert(g.Bind(t).Header(arg.Head).Delete(arg.URL), arg)
	})
}
func (g Matrix) postTest(arg MatrixTestArg) {
	g.t.Run(arg.Name, func(t *testing.T) {
		g.wantAssert(g.Bind(t).Header(arg.Head).Post(arg.URL, arg.Body), arg)
	})
}
func (g Matrix) putTest(arg MatrixTestArg) {
	g.t.Run(arg.Name, func(t *testing.T) {
		g.wantAssert(g.Bind(t).Header(arg.Head).Put(arg.URL, arg.Body), arg)
	})
}

// Test 进行矩阵测试
func (g Matrix) Test(args []MatrixTestArg) {
	for _, tt := range args {
		switch tt.Method {
		case http.MethodGet:
			g.getTest(tt)
		case http.MethodPost:
			g.postTest(tt)
		case http.MethodPut:
			g.putTest(tt)
		case http.MethodDelete:
			g.deleteTest(tt)
		default:
			g.t.Run(tt.Name, func(t *testing.T) {
				t.Errorf("没有实现%v方法测试", tt.Method)
			})
		}

	}
}

// GetTest 进行GET矩阵测试
func (g Matrix) GetTest(args []MatrixTestArg) {
	for _, tt := range args {
		g.getTest(tt)
	}
}

// PostTest 进行POST矩阵测试
func (g Matrix) PostTest(args []MatrixTestArg) {
	for _, tt := range args {
		g.postTest(tt)
	}
}

// PutTest 进行PUT矩阵测试
func (g Matrix) PutTest(args []MatrixTestArg) {
	for _, tt := range args {
		g.putTest(tt)
	}
}

// DeleteTest 进行DELETE矩阵测试
func (g Matrix) DeleteTest(args []MatrixTestArg) {
	for _, tt := range args {
		g.deleteTest(tt)
	}
}
