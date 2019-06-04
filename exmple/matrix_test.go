package server

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
