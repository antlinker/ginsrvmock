# gin 服务mock工具

## 获取

```bash

go get -u -v github.com/antlinker/ginsrvmock

```

## 使用例子

### 矩阵测试

[exmple源码](exmple/matrix_test.go)

``` golang
import (
	"fmt"
	"testing"

	"github.com/antlinker/ginsrvmock"
	"github.com/gin-gonic/gin"
)

func Test_Matrix_Demo(t *testing.T) {
	ginsrvmock.New(func(route *gin.Engine) {
		route.GET("/demo/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			if id == "world" {
				ctx.JSON(200, map[string]interface{}{
					"status":  0,
					"message": fmt.Sprintf("hello %v", id),
				})
				return
			}
			ctx.JSON(500, map[string]interface{}{
				"status":  1,
				"message": fmt.Sprintf("error %v", id),
			})

		})

	}).Bind(t).Matrix().GetTest([]ginsrvmock.MatrixTestArg{
		ginsrvmock.NewGetArg("true 1", "/demo/world", 200, map[string]interface{}{
			"status":  0,
			"message": "hello world",
		}),
		ginsrvmock.NewGetArg("error 1", "/demo/1", 500, map[string]interface{}{
			"status":  1,
			"message": "error 1",
		}),
	})
}

```
### 通用测试

[exmple源码](exmple/common_test.go)

```golang
import (
	"fmt"
	"testing"

	"github.com/antlinker/ginsrvmock"
	"github.com/gin-gonic/gin"
)

func Test_Common_Demo(t *testing.T) {
	ginmock := ginsrvmock.New(func(route *gin.Engine) {
		route.GET("/demo/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			if id == "world" {
				ctx.JSON(200, map[string]interface{}{
					"status":  0,
					"message": fmt.Sprintf("hello %v", id),
				})
				return
			}
			ctx.JSON(500, map[string]interface{}{
				"status":  1,
				"message": fmt.Sprintf("error %v", id),
			})

		})

	}).Bind(t)
	ginmock.Get("/demo/world").JSON().Assert(200, map[string]interface{}{
		"status":  0,
		"message": "hello world",
	})
	ginmock.Get("/demo/1").JSON().Assert(500, map[string]interface{}{
		"status":  1,
		"message": "error 1",
	})
}

```

## 适用范围

目前适用于gin服务，仅支持请求body是json格式，响应body也是json格式的请求.

支持　get post put delete 请求测试

支持http状态码断言

支持响应body体的断言

支持头信息的断言，头信息断言目前仅支持对想要断言的头信息进行完全匹配。

## 支持计划

1. 支持`application/x-www-form-urlencoded`格式
1. 支持`multipart/form-data`格式
1. 支持更复杂的断言模式，例如，可以部分匹配，目前主要是完全匹配
